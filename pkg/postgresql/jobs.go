package postgresql

import (
	"fmt"

	"github.com/diazharizky/scheduler/internal/definitions"
)

const (
	table         = "jobs"
	fields        = "job_id, entry_id, name, schedule, action, live, status"
	statusRunning = "running"
)

// GetRunningJobs func
func (p *PGInstance) GetRunningJobs() (jobs []definitions.Job, err error) {
	query := fmt.Sprintf(`SELECT %s FROM %s WHERE status = $1`, fields, table)
	err = p.Conn.Select(&jobs, query, statusRunning)

	return
}

// GetJob func
func (p *PGInstance) GetJob(jobID int64) (job definitions.Job, err error) {
	query := fmt.Sprintf(`SELECT %s FROM %s WHERE job_id = $1`, fields, table)
	err = p.Conn.Get(&job, query, jobID)

	return
}

// CreateJob func
func (p *PGInstance) CreateJob(job *definitions.Job) (entryID int64, err error) {
	query := fmt.Sprintf(`INSERT INTO %s (entry_id, name, schedule, action, live, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING job_id`, table)
	trx := p.Conn.MustBegin()
	err = trx.QueryRow(query, job.EntryID, job.Name, job.Schedule, job.Action, job.Live, job.Status).Scan(&entryID)
	trx.Commit()

	return
}

// UpdateJobStatus func
func (p *PGInstance) UpdateJobStatus(status string, jobID int64) (err error) {
	query := fmt.Sprintf(`UPDATE %s SET status = $1 WHERE job_id = $2`, table)
	trx := p.Conn.MustBegin()
	_, err = trx.Exec(query, status, jobID)
	trx.Commit()

	return
}
