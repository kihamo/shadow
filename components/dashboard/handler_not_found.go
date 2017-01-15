package dashboard

import (
	"net/http"
)

type NotFoundHandler struct {
	Handler
}

func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	h.Render(r.Context(), "dashboard", "404", nil)
}