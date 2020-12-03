package httphandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/diazharizky/scheduler/internal/definitions"
	"github.com/go-chi/chi"
)

const (
	statusRunning    = "running"
	statusTerminated = "terminated"
	paramJobID       = "job_id"
)

func (h *HTTPHandler) jobsRouter() (path string, r chi.Router) {
	path = "/jobs"

	r = chi.NewRouter()
	r.Get("/", h.getRunningJobs())
	r.Post("/", h.createJob())

	pathWithID := fmt.Sprintf(`/{%s}`, paramJobID)
	r.Get(pathWithID, h.getJob())
	r.Delete(pathWithID, h.stopJob())

	return
}

func (h *HTTPHandler) getRunningJobs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jobs, err := h.DB.GetRunningJobs()
		if err != nil {
			log.Print(err.Error())
		}

		json.NewEncoder(w).Encode(jobs)
	}
}

func (h *HTTPHandler) getJob() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jobIDParam := chi.URLParam(r, paramJobID)
		jobIDInt, err := strconv.ParseInt(jobIDParam, 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}

		job, err := h.DB.GetJob(jobIDInt)
		if err != nil {
			log.Print(err.Error())
		}

		json.NewEncoder(w).Encode(job)
	}
}

func (h *HTTPHandler) createJob() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var job definitions.Job
		json.NewDecoder(r.Body).Decode(&job)
		job.Status = statusRunning
		entryID, err := h.Cron.AddFunc(job.Schedule, func() {
			http.Get(job.Action)
			if job.Live == "once" {
				h.Cron.Remove(job.EntryID)
				err := h.DB.UpdateJobStatus(statusTerminated, job.ID)
				if err != nil {
					log.Print(err.Error())
				}
			}
		})
		if err != nil {
			log.Fatal(err.Error())
		}

		job.EntryID = entryID
		jobID, err := h.DB.CreateJob(&job)
		if err != nil {
			log.Fatal(err.Error())
		}

		job.ID = jobID

		json.NewEncoder(w).Encode(job)
	}
}

func (h *HTTPHandler) stopJob() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jobIDParam := chi.URLParam(r, paramJobID)
		jobIDInt, err := strconv.ParseInt(jobIDParam, 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}

		job, err := h.DB.GetJob(jobIDInt)
		if err != nil {
			log.Fatal(err.Error())
		}

		h.Cron.Remove(job.EntryID)
		err = h.DB.UpdateJobStatus(statusTerminated, jobIDInt)
		if err != nil {
			log.Fatal(err.Error())
		}

		json.NewEncoder(w).Encode(fmt.Sprintf("Job ID %d terminatted", jobIDInt))
	}
}
