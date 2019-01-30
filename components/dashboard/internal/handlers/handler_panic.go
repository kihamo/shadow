package handlers

import (
	"fmt"
	"net/http"

	"github.com/kihamo/shadow/components/dashboard"
)

type PanicHandler struct {
	dashboard.Handler
}

func (h *PanicHandler) ServeHTTP(w *dashboard.Response, r *dashboard.Request) {
	e := dashboard.PanicFromContext(r.Context())
	fields := map[string]interface{}{
		"error": fmt.Sprintf("%v", e.Error),
		"stack": string(e.Stack),
		"file":  e.File,
		"line":  e.Line,
	}

	w.WriteHeader(http.StatusInternalServerError)

	ctx := dashboard.ContextWithTemplateNamespace(r.Context(), dashboard.ComponentName)

	h.RenderLayout(ctx, "500", "simple", map[string]interface{}{
		"panic": fields,
	})

	h.Logger().Error("Internal server error",
		"error", fields["error"],
		"stack", fields["stack"],
		"file", fields["file"],
		"line", fields["line"],
	)
}
