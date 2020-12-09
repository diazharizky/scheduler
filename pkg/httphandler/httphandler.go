package httphandler

import (
	"net/http"

	"github.com/diazharizky/scheduler/internal/definitions"
	"github.com/go-chi/chi"
	"github.com/robfig/cron/v3"
)

// HTTPHandler contains mountable http handler
type HTTPHandler struct {
	Middlewares []func(next http.Handler) http.Handler
	Cron        *cron.Cron
	DB          definitions.DBAdapter
}

// Handler returns mountable http handler
func (h *HTTPHandler) Handler() (r *chi.Mux) {
	r = chi.NewRouter()
	for _, m := range h.Middlewares {
		r.Use(m)
	}
	r.Mount(h.jobsRouter())

	return r
}
