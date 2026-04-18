ALTER TABLE tasks
ADD COLUMN scheduled_at TIMESTAMPTZ,
ADD COLUMN parent_id BIGINT REFERENCES tasks(id) ON DELETE CASCADE,
ADD COLUMN recurrence_type TEXT,
ADD COLUMN recurrence_params JSONB;if matched {
    scheduledTime := date
    model := &taskdomain.Task{
        Title:       parent.Title,
        Description: parent.Description,
        Status:      taskdomain.StatusNew,
        ScheduledAt: &scheduledTime,
        ParentID:    &parent.ID,
        // time.Now().UTC() yerine halihazırda metodun başında çekilen 'now' değişkeni kullanıldı.
        CreatedAt:   now,
        UpdatedAt:   now,
    }
    _, err := s.repo.Create(ctx, model)
    if err != nil {
        return err
    }
}