package definitions

import "github.com/robfig/cron/v3"

// Job struct
type Job struct {
	ID       int64        `json:"job_id,omitempty" db:"job_id"`
	EntryID  cron.EntryID `json:"entry_id,omitempty" db:"entry_id"`
	Name     string       `json:"name" db:"name"`
	Schedule string       `json:"schedule" db:"schedule"`
	Action   string       `json:"action" db:"action"`
	Live     string       `json:"live,omitempty" db:"live"`
	Status   string       `json:"status,omitempty" db:"status"`
}

// DBAdapter interface for any database instance
type DBAdapter interface {
	MigrateUp() error
	MigrateDown() error
	Open()
	GetRunningJobs() ([]Job, error)
	CreateJob(job *Job) (int64, error)
	UpdateJobStatus(status string, jobID int64) error
}
