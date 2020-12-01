package postgresql

import (
	"github.com/diazharizky/scheduler/internal/definitions"
)

const table = "jobs"

// CreateSchedule func
func (p *PGInstance) CreateSchedule(job *definitions.Job) {
	tx := p.Conn.MustBegin()
	query := `INSERT INTO ` + table + ` (entry_id, name, schedule, action, live, status) VALUES ($1, $2, $3, $4, $5, $6)`
	tx.MustExec(query, job.EntryID, job.Name, job.Schedule, job.Action, job.Live, job.Status)
	tx.Commit()
}
