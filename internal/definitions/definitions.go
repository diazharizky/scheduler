package definitions

import "github.com/robfig/cron/v3"

// Job represent the routine/cron to be executed
//
// swagger:model
type Job struct {
	ID      int64        `json:"job_id,omitempty" db:"job_id"`
	EntryID cron.EntryID `json:"entry_id,omitempty" db:"entry_id"`
	Status  string       `json:"status,omitempty" db:"status"`

	// The name of routine/cron.
	// required: true
	Name string `json:"name" db:"name"`

	// The time schedule to be registered, mimics cron formatting. Read https://godoc.org/github.com/robfig/cron, for the reference.
	// required: true
	Schedule string `json:"schedule" db:"schedule"`

	// The action to be taken, so far only HTTP GET method is allowed.
	// required: true
	Action string `json:"action" db:"action"`

	// It defines how the service maintains the job, default to "once" if empty.
	// required: false
	Live string `json:"live,omitempty" db:"live"`
}

// DBAdapter interface for any database instance
type DBAdapter interface {
	MigrateUp() error
	MigrateDown() error
	Open()
	GetRunningJobs() ([]Job, error)
	GetJob(jobID int64) (Job, error)
	CreateJob(job *Job) (int64, error)
	UpdateJobStatus(status string, jobID int64) error
}
