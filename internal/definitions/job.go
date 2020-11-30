package definitions

// Job struct
type Job struct {
	Name     string `json:"name"`
	Schedule string `json:"schedule"`
	Action   string `json:"action"`
	Status   string `json:"status"`
}
