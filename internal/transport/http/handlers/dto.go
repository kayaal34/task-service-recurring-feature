package handlers

import (
	"time"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type taskMutationDTO struct {
	Title            string            `json:"title"`
	Description      string            `json:"description"`
	Status           taskdomain.Status `json:"status"`
	ScheduledAt      *time.Time        `json:"scheduled_at,omitempty"`
	ParentID         *int64            `json:"parent_id,omitempty"`
	RecurrenceType   *string           `json:"recurrence_type,omitempty"`
	RecurrenceParams *string           `json:"recurrence_params,omitempty"`
}

type taskDTO struct {
	ID               int64             `json:"id"`
	Title            string            `json:"title"`
	Description      string            `json:"description"`
	Status           taskdomain.Status `json:"status"`
	ScheduledAt      *time.Time        `json:"scheduled_at,omitempty"`
	ParentID         *int64            `json:"parent_id,omitempty"`
	RecurrenceType   *string           `json:"recurrence_type,omitempty"`
	RecurrenceParams *string           `json:"recurrence_params,omitempty"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
}

func newTaskDTO(task *taskdomain.Task) taskDTO {
	return taskDTO{
		ID:               task.ID,
		Title:            task.Title,
		Description:      task.Description,
		Status:           task.Status,
		ScheduledAt:      task.ScheduledAt,
		ParentID:         task.ParentID,
		RecurrenceType:   task.RecurrenceType,
		RecurrenceParams: task.RecurrenceParams,
		CreatedAt:        task.CreatedAt,
		UpdatedAt:        task.UpdatedAt,
	}
}
