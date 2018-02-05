package internal

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/dashboard/internal/handlers"
)

func (c *Component) DashboardTemplates() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "templates",
	}
}

func (c *Component) DashboardMenu() dashboard.Menu {
	routes := c.DashboardRoutes()

	return dashboard.NewMenuWithRoute(
		"Dashboard",
		routes[9],
		"dashboard",
		[]dashboard.Menu{
			dashboard.NewMenuWithRoute("Components", routes[8], "", nil, nil),
			dashboard.NewMenuWithRoute("Environment", routes[2], "", nil, nil),
			dashboard.NewMenuWithRoute("Bindata", routes[1], "", nil, nil),
			dashboard.NewMenuWithRoute("Routing", routes[3], "", nil, nil),
		},
		nil,
	)
}

func (c *Component) DashboardRoutes() []dashboard.Route {
	if c.routes == nil {
		c.routes = []dashboard.Route{
			dashboard.NewRoute(
				c.Name(),
				[]string{http.MethodGet},
				"/"+c.Name()+"/assets/*filepath",
				&assetfs.AssetFS{
					Asset:     Asset,
					AssetDir:  AssetDir,
					AssetInfo: AssetInfo,
					Prefix:    "assets",
				},
				"",
				false),
			dashboard.NewRoute(
				c.Name(),
				[]string{http.MethodGet},
				"/"+c.Name()+"/bindata",
				&handlers.BindataHandler{
					Application: c.application,
				},
				"",
				true),
			dashboard.NewRoute(
				c.Name(),
				[]string{http.MethodGet},
				"/"+c.Name()+"/environment",
				&handlers.EnvironmentHandler{},
				"",
				true),
			dashboard.NewRoute(
				c.Name(),
				[]string{http.MethodGet},
				"/"+c.Name()+"/routing",
				&handlers.RoutingHandler{},
				"",
				true),
			dashboard.NewRoute(
				c.Name(),
				[]string{http.MethodGet, http.MethodPost},
				dashboard.AuthPath+"/:provider/callback",
				&handlers.AuthHandler{
					IsCallback: true,
				},
				"",
				false),
			dashboard.NewRoute(
				c.Name(),
				[]string{http.MethodGet, http.MethodPost},
				dashboard.AuthPath+"/:provider",
				&handlers.AuthHandler{},
				"",
				false),
			dashboard.NewRoute(
				c.Name(),
				[]string{http.MethodGet},
				dashboard.AuthPath,
				&handlers.AuthHandler{},
				"",
				false),
			dashboard.NewRoute(
				c.Name(),
				[]string{http.MethodGet},
				"/"+c.Name()+"/logout",
				&handlers.LogoutHandler{},
				"",
				true),
		}

		componentsHandler := &handlers.ComponentsHandler{
			Application: c.application,
		}

		c.routes = append(c.routes, []dashboard.Route{
			dashboard.NewRoute(
				c.Name(),
				[]string{http.MethodGet},
				"/"+c.Name()+"/components",
				componentsHandler,
				"",
				true),
			dashboard.NewRoute(
				c.Name(),
				[]string{http.MethodGet},
				"/"+c.Name()+"/",
				componentsHandler,
				"",
				true),
		}...)
	}

	return c.routes
}
