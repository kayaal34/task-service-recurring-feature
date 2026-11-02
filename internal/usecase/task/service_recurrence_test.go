package task

import (
	"context"
	"testing"
	"time"
)

func createRecurring(t *testing.T, start time.Time, rtype string, params *string) []time.Time {
	t.Helper()
	repo := &fakeRepository{}

	input := CreateInput{Title: "task", ScheduledAt: &start, RecurrenceType: &rtype, RecurrenceParams: params}
	if _, err := NewService(repo).Create(context.Background(), input); err != nil {
		t.Fatalf("Create returned an error: %v", err)
	}

	dates := make([]time.Time, 0, len(repo.created))
	for _, child := range repo.created[1:] {
		if child.ParentID == nil || *child.ParentID != repo.created[0].ID {
			t.Fatalf("child task has parent %v", child.ParentID)
		}
		dates = append(dates, *child.ScheduledAt)
	}
	return dates
}

func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 9, 0, 0, 0, time.UTC)
}

func TestDailyCreatesThirtyChildren(t *testing.T) {
	dates := createRecurring(t, date(2026, time.March, 1), "daily", nil)

	if len(dates) != 30 {
		t.Fatalf("got %d children", len(dates))
	}
	if !dates[0].Equal(date(2026, time.March, 2)) {
		t.Errorf("first child = %v", dates[0])
	}
}

func TestMonthlyCreatesNextOccurrence(t *testing.T) {
	cases := map[string]struct{ start, next time.Time }{
		"30 day month":     {date(2026, time.April, 15), date(2026, time.May, 15)},
		"31 day month":     {date(2026, time.January, 1), date(2026, time.February, 1)},
		"31 day month end": {date(2026, time.March, 15), date(2026, time.April, 15)},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			dates := createRecurring(t, tc.start, "monthly", nil)

			if len(dates) != 1 || !dates[0].Equal(tc.next) {
				t.Fatalf("got %v, want %v", dates, tc.next)
			}
		})
	}
}

func TestEvenDaysOnly(t *testing.T) {
	even := "even"

	for _, day := range createRecurring(t, date(2026, time.March, 1), "even_odd", &even) {
		if day.Day()%2 != 0 {
			t.Errorf("child on odd day %v", day)
		}
	}
}

func TestExactDates(t *testing.T) {
	params := "2026-03-05,2026-03-20"

	dates := createRecurring(t, date(2026, time.March, 1), "exact_dates", &params)

	if len(dates) != 2 || dates[0].Day() != 5 || dates[1].Day() != 20 {
		t.Fatalf("got %v", dates)
	}
}
