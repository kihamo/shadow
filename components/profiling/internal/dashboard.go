package internal

import (
	"net/http"
	pprofHandlers "net/http/pprof"
	"runtime/pprof"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/profiling/internal/handlers"
)

func (c *Component) GetTemplates() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "templates",
	}
}

func (c *Component) GetDashboardMenu() dashboard.Menu {
	routes := c.GetDashboardRoutes()

	show := func(r *dashboard.Request) bool {
		return r.Config().GetBool(config.ConfigDebug)
	}

	return dashboard.NewMenuItemWithUrl(
		"Profiling",
		"",
		"terminal",
		[]dashboard.Menu{
			dashboard.NewMenuItemWithRoute("Traces", routes[1], "", nil, show),
			dashboard.NewMenuItemWithRoute("Pprof", routes[7], "", nil, show),
			dashboard.NewMenuItemWithRoute("Expvar", routes[2], "", nil, show),
		},
		show,
	)
}

func (c *Component) GetDashboardRoutes() []dashboard.Route {
	if c.routes == nil {
		c.routes = []dashboard.Route{
			dashboard.NewRouteItem(
				c.GetName(),
				[]string{http.MethodGet},
				"/"+c.GetName()+"/assets/*filepath",
				&assetfs.AssetFS{
					Asset:     Asset,
					AssetDir:  AssetDir,
					AssetInfo: AssetInfo,
					Prefix:    "assets",
				},
				"",
				false),
			dashboard.NewRouteItem(
				c.GetName(),
				[]string{http.MethodGet, http.MethodPost},
				"/"+c.GetName()+"/trace/",
				&handlers.TraceHandler{},
				"",
				true),
			dashboard.NewRouteItem(
				c.GetName(),
				[]string{http.MethodGet},
				"/debug/vars/",
				&handlers.ExpvarHandler{},
				"",
				false),
			dashboard.NewRouteItem(
				c.GetName(),
				[]string{http.MethodGet},
				"/debug/pprof/cmdline",
				&handlers.DebugHandler{
					HandlerFunc: pprofHandlers.Cmdline,
				},
				"",
				false),
			dashboard.NewRouteItem(
				c.GetName(),
				[]string{http.MethodGet},
				"/debug/pprof/profile",
				&handlers.DebugHandler{
					HandlerFunc: pprofHandlers.Profile,
				},
				"",
				false),
			dashboard.NewRouteItem(
				c.GetName(),
				[]string{http.MethodGet},
				"/debug/pprof/symbol",
				&handlers.DebugHandler{
					HandlerFunc: pprofHandlers.Symbol,
				},
				"",
				false),
			dashboard.NewRouteItem(
				c.GetName(),
				[]string{http.MethodGet},
				"/debug/pprof/trace",
				&handlers.DebugHandler{
					HandlerFunc: pprofHandlers.Trace,
				},
				"",
				false),
			dashboard.NewRouteItem(
				c.GetName(),
				[]string{http.MethodGet},
				"/debug/pprof/",
				&handlers.DebugHandler{
					HandlerFunc: pprofHandlers.Index,
				},
				"",
				false),
		}

		for _, profile := range pprof.Profiles() {
			c.routes = append(c.routes, dashboard.NewRouteItem(
				c.GetName(),
				[]string{http.MethodGet},
				"/debug/pprof/"+profile.Name(),
				&handlers.DebugHandler{
					HandlerFunc: pprofHandlers.Index,
				},
				"",
				false))
		}
	}

	return c.routes
}
