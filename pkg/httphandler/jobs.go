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
	path             = "/jobs"
	paramJobID       = "job_id"
	statusRunning    = "running"
	statusTerminated = "terminated"
	liveOnce         = "once"
)

func (h *HTTPHandler) jobsRouter() (string, chi.Router) {
	r := chi.NewRouter()
	r.Get("/", h.getRunningJobs())
	r.Post("/", h.createJob())

	pathWithID := fmt.Sprintf(`/{%s}`, paramJobID)
	r.Get(pathWithID, h.getJob())
	r.Delete(pathWithID, h.stopJob())

	return path, r
}

// swagger:operation GET /jobs jobs
//
// Get current running jobs.
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: Successful operation
func (h *HTTPHandler) getRunningJobs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jobs, err := h.DB.GetRunningJobs()
		if err != nil {
			log.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jobs)
	}
}

// swagger:operation GET /jobs/{job_id} jobs
//
// Get specific job.
//
// ---
// produces:
// - application/json
// parameters:
// - name: job_id
//   in: path
//   description: Retrieved from `/jobs`
//   required: true
//   type: integer
//   format: int64
// responses:
//   '200':
//     description: Successful operation
func (h *HTTPHandler) getJob() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jobIDParam := chi.URLParam(r, paramJobID)
		jobIDInt, err := strconv.ParseInt(jobIDParam, 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}

		job, err := h.DB.GetJob(jobIDInt)
		if err != nil {
			log.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(job)
	}
}

// swagger:operation POST /jobs jobs
//
// Add a new cron job.
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - in: body
//   name: body
//   description: Job payload.
//   required: true
//   schema:
//     $ref: "#/definitions/JobPayload"
// responses:
//   '200':
//     description: "Successful operation"
func (h *HTTPHandler) createJob() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var job definitions.Job
		json.NewDecoder(r.Body).Decode(&job)
		job.Status = statusRunning
		entryID, err := h.Cron.AddFunc(job.Schedule, func() {
			http.Get(job.Action)
			if job.Live == liveOnce {
				h.Cron.Remove(job.EntryID)
				err := h.DB.UpdateJobStatus(statusTerminated, job.ID)
				if err != nil {
					log.Println(err.Error())
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

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(job)
	}
}

// swagger:operation DELETE /jobs/{job_id} jobs
//
// Stop specific running/terminated job
//
// ---
// produces:
// - text/plain
// parameters:
// - name: job_id
//   in: path
//   description: Retrieved from `/jobs`
//   required: true
//   type: integer
//   format: int64
// responses:
//   '200':
//     description: Successful operation
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
