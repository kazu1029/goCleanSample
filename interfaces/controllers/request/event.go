package request

type Event struct {
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	UserID      int    `form:"user_id" json:"user_id" binding:"required"`
}
