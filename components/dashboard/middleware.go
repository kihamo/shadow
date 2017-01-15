package dashboard

import (
	"context"
	"net/http"
	"time"

	"github.com/justinas/alice"
)

func ContextMiddleware(c *Component) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), ConfigContextKey, c.config)
			ctx = context.WithValue(ctx, LoggerContextKey, c.logger)
			ctx = context.WithValue(ctx, RenderContextKey, c.renderer)
			ctx = context.WithValue(ctx, RequestContextKey, r)
			ctx = context.WithValue(ctx, ResponseContextKey, w)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func BasicAuthMiddleware(_ *Component) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			config := ConfigFromContext(r.Context())

			checkUsername := config.GetString(ConfigFrontendAuthUser)
			if checkUsername == "" {
				next.ServeHTTP(w, r)
				return
			}

			checkPassword := config.GetString(ConfigFrontendAuthPassword)

			username, password, ok := r.BasicAuth()
			if ok && checkUsername == username && checkPassword == password {
				next.ServeHTTP(w, r)
				return
			}

			w.Header().Set("WWW-Authenticate", `Basic realm="Shadow Security Zone"`)
			w.WriteHeader(http.StatusUnauthorized)
		})
	}
}

func LoggerMiddleware(c *Component) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			writer := &ResponseWriter{
				ResponseWriter: w,
			}

			next.ServeHTTP(writer, r)

			fields := map[string]interface{}{
				"remote-addr":    r.RemoteAddr,
				"method":         r.Method,
				"request-uri":    r.RequestURI,
				"prote":          r.Proto,
				"code":           writer.GetStatusCode(),
				"content-length": r.ContentLength,
				"referer":        r.Referer(),
				"user-agent":     r.UserAgent(),
			}

			logger := LoggerFromContext(r.Context())

			if writer.GetStatusCode()/100 == 5 {
				logger.Error(http.StatusText(writer.GetStatusCode()), fields)
			} else {
				logger.Info(http.StatusText(writer.GetStatusCode()), fields)
			}
		})
	}
}

func MetricsMiddleware(_ *Component) alice.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()

			next.ServeHTTP(w, r)

			if metricHandlerExecuteTime != nil {
				metricHandlerExecuteTime.ObserveDurationByTime(now)
			}
		})
	}
}