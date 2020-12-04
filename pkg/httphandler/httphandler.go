package httphandler

import (
	"log"
	"net/http"
	"os"

	"github.com/diazharizky/scheduler/internal/definitions"
	"github.com/ghodss/yaml"
	"github.com/go-chi/chi"
	"github.com/gobuffalo/packr/v2"
	"github.com/robfig/cron/v3"
	httpSwagger "github.com/swaggo/http-swagger"
)

// HTTPHandler contains mountable http handler
type HTTPHandler struct {
	Cron *cron.Cron
	DB   definitions.DBAdapter
}

// Handler returns mountable http handler
func (h *HTTPHandler) Handler() (r *chi.Mux) {
	r = chi.NewRouter()

	swaggerSourcePath := "/swagger/source"
	r.Get(swaggerSourcePath, swaggerSource())
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerSourcePath),
	))
	r.Mount(h.jobsRouter())

	return r
}

func swaggerSource() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pwd, _ := os.Getwd()
		configPath := pwd + "/configs/"
		box := packr.New("configs", configPath)
		yamlFile, err := box.FindString("swagger.yml")
		if err != nil {
			log.Fatal(err)
		}

		swaggerDocs, err := yaml.YAMLToJSON([]byte(yamlFile))
		if err != nil {
			log.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(swaggerDocs)
	}
}
