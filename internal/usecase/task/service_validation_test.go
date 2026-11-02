package task

import (
	"context"
	"errors"
	"testing"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type fakeRepository struct {
	created []*taskdomain.Task
}

func (r *fakeRepository) Create(_ context.Context, task *taskdomain.Task) (*taskdomain.Task, error) {
	stored := *task
	stored.ID = int64(len(r.created) + 1)
	r.created = append(r.created, &stored)
	return &stored, nil
}

func (r *fakeRepository) GetByID(context.Context, int64) (*taskdomain.Task, error) {
	return nil, nil
}

func (r *fakeRepository) Update(_ context.Context, task *taskdomain.Task) (*taskdomain.Task, error) {
	return task, nil
}

func (r *fakeRepository) Delete(context.Context, int64) error {
	return nil
}

func (r *fakeRepository) List(context.Context) ([]taskdomain.Task, error) {
	return nil, nil
}

func ptr(value string) *string {
	return &value
}

func TestCreateTrimsAndDefaultsStatus(t *testing.T) {
	repo := &fakeRepository{}

	created, err := NewService(repo).Create(context.Background(), CreateInput{Title: "  write report  "})
	if err != nil {
		t.Fatalf("Create returned an error: %v", err)
	}
	if created.Title != "write report" {
		t.Errorf("title = %q", created.Title)
	}
	if created.Status != taskdomain.StatusNew {
		t.Errorf("status = %q", created.Status)
	}
}

func TestCreateRejectsInvalidInput(t *testing.T) {
	cases := map[string]CreateInput{
		"empty title":             {Title: "   "},
		"unknown status":          {Title: "task", Status: "paused"},
		"unknown recurrence":      {Title: "task", RecurrenceType: ptr("weekly")},
		"even_odd without param":  {Title: "task", RecurrenceType: ptr("even_odd")},
		"even_odd wrong param":    {Title: "task", RecurrenceType: ptr("even_odd"), RecurrenceParams: ptr("both")},
		"exact_dates no dates":    {Title: "task", RecurrenceType: ptr("exact_dates")},
		"exact_dates blank dates": {Title: "task", RecurrenceType: ptr("exact_dates"), RecurrenceParams: ptr("  ")},
	}

	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			repo := &fakeRepository{}

			_, err := NewService(repo).Create(context.Background(), input)
			if !errors.Is(err, ErrInvalidInput) {
				t.Fatalf("error = %v, want ErrInvalidInput", err)
			}
			if len(repo.created) != 0 {
				t.Errorf("%d tasks were stored", len(repo.created))
			}
		})
	}
}

func TestUpdateRejectsUnknownRecurrence(t *testing.T) {
	input := UpdateInput{Title: "task", Status: taskdomain.StatusNew, RecurrenceType: ptr("weekly")}

	_, err := NewService(&fakeRepository{}).Update(context.Background(), 1, input)
	if !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("error = %v, want ErrInvalidInput", err)
	}
}

func TestIDsMustBePositive(t *testing.T) {
	service := NewService(&fakeRepository{})

	if _, err := service.GetByID(context.Background(), 0); !errors.Is(err, ErrInvalidInput) {
		t.Errorf("GetByID error = %v", err)
	}
	if err := service.Delete(context.Background(), -1); !errors.Is(err, ErrInvalidInput) {
		t.Errorf("Delete error = %v", err)
	}
}
