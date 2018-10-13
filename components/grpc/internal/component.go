package internal

import (
	"net"
	"os"
	"sync"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/grpc"
	"github.com/kihamo/shadow/components/grpc/server"
	"github.com/kihamo/shadow/components/grpc/stats"
	"github.com/kihamo/shadow/components/i18n"
	"github.com/kihamo/shadow/components/logger"
	"github.com/kihamo/shadow/components/metrics"
	g "google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	s "google.golang.org/grpc/stats"
)

type Component struct {
	application shadow.Application
	config      config.Component
	logger      logger.Logger
	server      *g.Server
	routes      []dashboard.Route
}

func (c *Component) Name() string {
	return grpc.ComponentName
}

func (c *Component) Version() string {
	return grpc.ComponentVersion + "/" + g.Version
}

func (c *Component) Dependencies() []shadow.Dependency {
	return []shadow.Dependency{
		{
			Name:     config.ComponentName,
			Required: true,
		},
		{
			Name: i18n.ComponentName,
		},
		{
			Name: logger.ComponentName,
		},
		{
			Name: metrics.ComponentName,
		},
	}
}

func (c *Component) Init(a shadow.Application) error {
	c.application = a
	c.config = a.GetComponent(config.ComponentName).(config.Component)

	return nil
}

func (c *Component) Run(wg *sync.WaitGroup) error {
	c.logger = logger.NewOrNop(c.Name(), c.application)
	grpclog.SetLoggerV2(grpc.NewLogger(c.logger))

	components, err := c.application.GetComponents()
	if err != nil {
		return err
	}

	unaryInterceptors := make([]g.UnaryServerInterceptor, 0, 0)
	streamInterceptors := make([]g.StreamServerInterceptor, 0, 0)
	statsHandlers := []s.Handler{
		stats.NewContextHandler(c.config),
	}

	for _, cmp := range components {
		if cmpUnaryServerInterceptors, ok := cmp.(grpc.HasUnaryServerInterceptors); ok {
			unaryInterceptors = append(unaryInterceptors, cmpUnaryServerInterceptors.GrpcUnaryServerInterceptors()...)
		}

		if cmpStreamServerInterceptors, ok := cmp.(grpc.HasStreamServerInterceptors); ok {
			streamInterceptors = append(streamInterceptors, cmpStreamServerInterceptors.GrpcStreamServerInterceptors()...)
		}

		if cmpStatsHandlers, ok := cmp.(grpc.HasStatsHandlers); ok {
			statsHandlers = append(statsHandlers, cmpStatsHandlers.GrpcStatsHandlers()...)
		}
	}

	c.server = server.NewDefaultServerWithCustomOptions(unaryInterceptors, streamInterceptors, statsHandlers)

	for _, cmp := range components {
		if cmpGrpc, ok := cmp.(grpc.HasGrpcServer); ok {
			cmpGrpc.RegisterGrpcServer(c.server)
		}
	}

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(c.server, healthServer)

	for service, info := range c.server.GetServiceInfo() {
		healthServer.SetServingStatus(service, grpc_health_v1.HealthCheckResponse_SERVING)

		for _, method := range info.Methods {
			c.logger.Debugf("Add method /%s/%s", service, method.Name)
		}
	}

	var wgListen sync.WaitGroup
	wgListen.Add(1)

	go func() {
		defer wg.Done()

		if c.config.Bool(grpc.ConfigReflectionEnabled) {
			reflection.Register(c.server)
		}

		addr := net.JoinHostPort(c.config.String(grpc.ConfigHost), c.config.String(grpc.ConfigPort))
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			c.logger.Fatalf("Failed to listen [%d]: %s\n", os.Getpid(), err.Error())
		}

		c.logger.Info("Running service", map[string]interface{}{
			"addr": addr,
			"pid":  os.Getpid(),
		})

		wgListen.Done()

		if err := c.server.Serve(lis); err != nil {
			c.logger.Fatalf("Failed to serve [%d]: %s\n", os.Getpid(), err.Error())
		}
	}()

	wgListen.Wait()
	return nil
}

func (c *Component) GetServiceInfo() map[string]g.ServiceInfo {
	if c.server == nil {
		return nil
	}

	return c.server.GetServiceInfo()
}
