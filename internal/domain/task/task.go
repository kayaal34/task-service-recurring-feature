package task

import "time"

type Status string

const (
	StatusNew        Status = "new"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID               int64      `json:"id"`
	Title            string     `json:"title"`
	Description      string     `json:"description"`
	Status           Status     `json:"status"`
	ScheduledAt      *time.Time `json:"scheduled_at,omitempty"`
	ParentID         *int64     `json:"parent_id,omitempty"`
	RecurrenceType   *string    `json:"recurrence_type,omitempty"`
	RecurrenceParams *string    `json:"recurrence_params,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

func (s Status) Valid() bool {
	switch s {
	case StatusNew, StatusInProgress, StatusDone:
		return true
	default:
		return false
	}
}
