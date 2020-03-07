package domain

type User struct {
	ID        int
	FirstName string
	LastName  string
	NickName  string
	Email     string
}

func NewUser(firstName, lastName, nickName, email string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		NickName:  nickName,
		Email:     email,
	}
}
