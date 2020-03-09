package iRepository

import "github.com/kazu1029/goCleanSample/domain"

type EventRepository struct {
	SqlHandler
}

func (repo *EventRepository) Store(e *domain.Event) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO events (user_id, title, description) VALUES (?,?,?)",
		e.UserID, e.Title, e.Description,
	)
	if err != nil {
		return
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}

	id = int(id64)
	return
}

func (repo *EventRepository) FindById(identifier int) (event domain.Event, err error) {
	row, err := repo.Query(
		"SELECT id, user_id, title, description FROM events WHERE id = ?",
		identifier,
	)
	defer row.Close()
	if err != nil {
		return
	}

	if err = row.Scan(&event.ID, &event.UserID, &event.Title, &event.Description); err != nil {
		return
	}
	return
}

func (repo *EventRepository) FindAll() (events []domain.Events, err error) {
	rows, err := repo.Query("SELECT id, user_id, title, description FROM events")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var event domain.Event
		if err := rows.Scan(&event.ID, &event.UserID, &event.Title, &event.Description); err != nil {
			continue
		}
		events = append(events, event)
	}

	return
}
