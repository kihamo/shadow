package grpc

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func HealthCheck(conn grpc.ClientConnInterface, service string) error {
	healthClient := grpc_health_v1.NewHealthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{
		Service: service,
	})
	if err != nil {
		return err
	}

	if response.Status != grpc_health_v1.HealthCheckResponse_SERVING {
		return errors.New("server is not healthy. Status is " + response.Status.String())
	}

	return nil
}
