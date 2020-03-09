package domain

type User struct {
	ID       int
	NickName string `validate:"required,gte=3,lte=20"`
	Email    string `validate:"required,email"`
}

func NewUser(nickName, email string) *User {
	return &User{
		NickName: nickName,
		Email:    email,
	}
}
