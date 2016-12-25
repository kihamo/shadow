package frontend

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"sync"

	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/resource/config"
	"github.com/kihamo/shadow/resource/logger"
	"github.com/kihamo/shadow/resource/template"
)

type FrontendMenu struct {
	Name    string
	Url     string
	Icon    string
	SubMenu []*FrontendMenu
}

type ServiceFrontendHandlers interface {
	SetFrontendHandlers(*Router)
}

type ServiceFrontendMenu interface {
	GetFrontendMenu() *FrontendMenu
}

type FrontendService struct {
	application *shadow.Application
	config      *config.Resource
	template    *template.Resource
	logger      logger.Logger
	router      *Router

	mutex     sync.RWMutex
	authToken string
}

func (c *FrontendService) GetName() string {
	return "frontend"
}

func (s *FrontendService) Init(a *shadow.Application) (err error) {
	s.application = a

	resourceTemplate, err := a.GetResource("template")
	if err != nil {
		return err
	}
	s.template = resourceTemplate.(*template.Resource)
	s.template.Globals["Menu"] = make([]*FrontendMenu, 0, len(s.application.GetServices()))

	resourceConfig, err := a.GetResource("config")
	if err != nil {
		return err
	}
	s.config = resourceConfig.(*config.Resource)

	s.template.Globals["AlertsEnabled"] = a.HasResource("alerts")

	return nil
}

func (s *FrontendService) Run(wg *sync.WaitGroup) error {
	s.logger = logger.NewOrNop(s.GetName(), s.application)

	s.generateAuthToken(s.config.GetString(ConfigFrontendAuthUser), s.config.GetString(ConfigFrontendAuthPassword))

	// скидывает mux по-умолчанию, так как pprof добавил свои хэндлеры
	http.DefaultServeMux = http.NewServeMux()

	s.router = NewRouter(s)
	s.router.SetPanicHandler(s, &PanicHandler{})
	s.router.SetNotAllowedHandler(s, &MethodNotAllowedHandler{})
	s.router.SetNotFoundHandler(s, &NotFoundHandler{})

	if s.config.GetBool(config.ConfigDebug) {
		s.router.Handler("GET", "/debug/pprof/cmdline", s.debugHandler(pprof.Cmdline))
		s.router.Handler("GET", "/debug/pprof/profile", s.debugHandler(pprof.Profile))
		s.router.Handler("GET", "/debug/pprof/symbol", s.debugHandler(pprof.Symbol))
		s.router.Handler("POST", "/debug/pprof/symbol", s.debugHandler(pprof.Symbol))
		s.router.Handler("GET", "/debug/pprof/block", s.debugHandler(pprof.Index))
		s.router.Handler("GET", "/debug/pprof/goroutine", s.debugHandler(pprof.Index))
		s.router.Handler("GET", "/debug/pprof/heap", s.debugHandler(pprof.Index))
		s.router.Handler("GET", "/debug/pprof/threadcreate", s.debugHandler(pprof.Index))
		s.router.Handler("GET", "/debug/pprof/", s.debugHandler(pprof.Index))
	}

	menus := make([]*FrontendMenu, 0, len(s.application.GetServices()))

	for _, service := range s.application.GetServices() {
		if serviceHandlers, ok := service.(ServiceFrontendHandlers); ok {
			serviceHandlers.SetFrontendHandlers(s.router)
		}

		if serviceMenu, ok := service.(ServiceFrontendMenu); ok {
			menu := serviceMenu.GetFrontendMenu()
			if menu != nil {
				if service == s {
					menus = append([]*FrontendMenu{menu}, menus...)
				} else {
					menus = append(menus, menu)
				}
			}
		}
	}

	s.template.Globals["Menu"] = menus

	go func(router *Router) {
		defer wg.Done()

		http.HandleFunc("/", func(out http.ResponseWriter, in *http.Request) {
			router.ServeHTTP(out, in)
		})

		// TODO: ssl
		addr := fmt.Sprintf("%s:%d", s.config.GetString(ConfigFrontendHost), s.config.GetInt(ConfigFrontendPort))

		s.logger.Info("Running service", map[string]interface{}{
			"addr": addr,
			"pid":  os.Getpid(),
		})

		if err := http.ListenAndServe(addr, s.router); err != nil {
			s.logger.Fatalf("Could not start frontend [%d]: %s\n", os.Getpid(), err.Error())
		}
	}(s.router)

	return nil
}

func (s *FrontendService) generateAuthToken(user, password string) {
	token := ""
	if user != "" {
		token = "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password))
	}

	s.mutex.Lock()
	s.authToken = token
	s.mutex.Unlock()
}

func (s *FrontendService) debugHandler(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if s.config.GetBool(config.ConfigDebug) {
			h.ServeHTTP(w, r)
		} else {
			s.router.NotFound.ServeHTTP(w, r)
		}
	})
}
