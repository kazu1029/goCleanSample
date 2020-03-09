package domain

type Event struct {
	ID          int
	UserID      int `validate:"required"`
	Title       string
	Description string `validate:"required,gte=8,lte=1000"`
}

func NewEvent(title, description string) *Event {
	return &Event{
		Title:       title,
		Description: description,
	}
}
