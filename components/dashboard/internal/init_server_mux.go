package internal

import (
	"net/http"

	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/dashboard/internal/handlers"
)

func (c *Component) initServeMux() error {
	// Special pages
	c.router.SetPanicHandler(handlers.NewPanicHandler(c))
	c.router.SetNotFoundHandler(handlers.NewNotFoundHandler(c))
	c.router.SetNotAllowedHandler(handlers.NewMethodNotAllowedHandler(c))

	// Middleware
	c.router.addMiddleware(ContextMiddleware(c.application, c.router, c.config, c.renderer, c.session))

	for _, component := range c.components {
		if componentRoute, ok := component.(dashboard.HasRoutes); ok {
			for _, route := range componentRoute.DashboardRoutes() {
				c.router.addRoute(NewRouteItem(route, component))
			}
		}

		if componentMiddleware, ok := component.(dashboard.HasServerMiddleware); ok {
			for _, middleware := range componentMiddleware.DashboardMiddleware() {
				c.router.addMiddleware(middleware)
			}
		}
	}

	// fixing middleware
	c.router.addMiddleware(AuthorizationMiddleware)

	// fixing routes
	startUrlRoute := NewRouteItem(dashboard.NewRoute("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, c.config.String(dashboard.ConfigStartURL), http.StatusMovedPermanently)
	})), c)

	c.router.addRoute(startUrlRoute)

	return nil
}
