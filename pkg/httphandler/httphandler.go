package httphandler

import (
	"log"
	"net/http"

	"github.com/diazharizky/scheduler/internal/definitions"
	"github.com/go-chi/chi"
	"github.com/gobuffalo/packr/v2"
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

	r.Get("/swagger.json", swaggerSource())
	r.Mount(h.jobsRouter())

	return r
}

func swaggerSource() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		box := packr.New("api", "../../api/swagger-spec/")
		swaggerSource, err := box.FindString("police.json")
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(swaggerSource))
	}
}
