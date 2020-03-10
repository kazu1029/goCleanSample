package domain

type Event struct {
	ID          int
	UserID      int `validate:"required"`
	Title       string
	Description string `validate:"required,gte=8,lte=1000"`
}

func NewEvent(user_id int, title, description string) *Event {
	return &Event{
		UserID:      user_id,
		Title:       title,
		Description: description,
	}
}
