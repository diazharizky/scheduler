package definitions

import "github.com/robfig/cron/v3"

// Job struct
type Job struct {
	EntryID  cron.EntryID `json:"entry_id"`
	Name     string       `json:"name"`
	Schedule string       `json:"schedule"`
	Action   string       `json:"action"`
	Live     string       `json:"live"`
	Status   string       `json:"status"`
}

// DBAdapter interface
type DBAdapter interface {
	Open() error
	CreateSchedule() error
}
