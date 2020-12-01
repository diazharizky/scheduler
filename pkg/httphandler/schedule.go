package httphandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/diazharizky/scheduler/internal/definitions"
	"github.com/go-chi/chi"
)

const pendingStatus = "pending"

func (h *HTTPHandler) scheduleRouter() (path string, r chi.Router) {
	path = "/schedules"

	r = chi.NewRouter()
	r.Post("/", h.createSchedule())

	return
}

func (h *HTTPHandler) createSchedule() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var job definitions.Job
		json.NewDecoder(r.Body).Decode(&job)
		job.Status = pendingStatus

		jobID, err := h.Cron.AddFunc(job.Schedule, func() {
			http.Get(job.Action)

			if job.Live == "once" {
				h.Cron.Remove(job.EntryID)
			}
		})
		if err != nil {
			log.Fatal(err.Error())
		}

		job.EntryID = jobID
		h.DB.CreateSchedule(&job)

		json.NewEncoder(w).Encode(job)
	}
}
