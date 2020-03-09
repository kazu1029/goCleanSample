package request

type User struct {
	NickName string `form:"nick_name" json:"nick_name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}
