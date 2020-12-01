package httphandler

import (
	"encoding/json"
	"fmt"
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
			fmt.Println("Every hour on the half hour")
			if job.Live == "once" {
				h.Cron.Remove(job.EntryID)
			}
		})
		if err != nil {
			log.Fatal(err)
		}
		job.EntryID = jobID
		json.NewEncoder(w).Encode(job)
	}
}

func (h *HTTPHandler) loadSchedules() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var jobs []definitions.Job
		json.NewEncoder(w).Encode(jobs)
	}
}