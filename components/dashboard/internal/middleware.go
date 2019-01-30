package internal

import (
	"net/http"

	"github.com/alexedwards/scs"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/dashboard/auth"
)

func ContextMiddleware(router *Router, renderer *Renderer, sessionManager *scs.Manager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			writer := dashboard.NewResponse(w)
			request := dashboard.NewRequest(r)
			route := dashboard.RouteFromContext(r.Context())

			ctx := dashboard.ContextWithRender(r.Context(), renderer)
			ctx = dashboard.ContextWithResponse(ctx, writer)
			ctx = dashboard.ContextWithRouter(ctx, router)
			ctx = dashboard.ContextWithSession(ctx, NewSession(sessionManager.Load(r), w))
			ctx = dashboard.ContextWithRequest(ctx, request)

			if route != nil {
				if routeItem, ok := route.(*RouteItem); ok {
					ctx = dashboard.ContextWithTemplateNamespace(ctx, routeItem.Component().Name())
				}
			}

			request = request.WithContext(ctx)
			next.ServeHTTP(writer, request.Original())
		})
	}
}

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := dashboard.RouteFromContext(r.Context())
		if route != nil && route.Auth() {
			request := dashboard.RequestFromContext(r.Context())
			if request == nil {
				panic("Request isn't set in context")
			}

			if len(auth.GetProviders()) > 0 && !request.User().IsAuthorized() {
				if !request.IsAjax() && request.IsGet() {
					if err := request.Session().PutString(dashboard.SessionLastURL, request.URL().Path); err != nil {
						panic(err.Error())
					}
				}

				http.Redirect(w, r, dashboard.AuthPath, http.StatusFound)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
