package httphandler

import (
	"github.com/diazharizky/scheduler/internal/definitions"
	"github.com/go-chi/chi"
	"github.com/robfig/cron/v3"
)

// HTTPHandler contains mountable http handler
type HTTPHandler struct {
	Cron *cron.Cron
	DB   definitions.DBAdapter
}

// Handler returns mountable http handler
func (h *HTTPHandler) Handler() (r *chi.Mux) {
	r = chi.NewRouter()

	r.Mount(h.jobsRouter())

	return r
}
