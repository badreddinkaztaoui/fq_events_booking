package repository

import (
	"context"
	"database/sql"

	"github.com/badreddinkaztaoui/fq_events_booking/internal/models"
)

type EventRepo struct {
	DB *sql.DB
}

func NewEventRepo(db *sql.DB) *EventRepo {
	return &EventRepo{DB: db}
}

func (repo *EventRepo) GetAll(ctx context.Context) ([]models.Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events`
	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event

	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (repo *EventRepo) Create(ctx context.Context, event *models.Event) error {
	query := `INSERT INTO events (name, description, location, date_time, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	row := repo.DB.QueryRowContext(ctx, query, event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	return row.Scan(&event.ID)
}

func (repo *EventRepo) GetByID(ctx context.Context, id int64) (*models.Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events WHERE id = $1`
	row := repo.DB.QueryRowContext(ctx, query, id)

	var event models.Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &event, nil
}

func (repo *EventRepo) Update(ctx context.Context, id int64, event *models.Event) error {
	query := `
        UPDATE events
        SET name = $1, description = $2, location = $3, date_time = $4
        WHERE id = $5
        RETURNING id, name, description, location, date_time, user_id
    `
	row := repo.DB.QueryRowContext(ctx, query, event.Name, event.Description, event.Location, event.DateTime, id)
	return row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
}

func (repo *EventRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM events WHERE id = $1`
	_, err := repo.DB.ExecContext(ctx, query, id)
	return err
}
