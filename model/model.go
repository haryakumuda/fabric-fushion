package model

type User struct {
	Id       int64
	Email    string
	Password string
	Role     string
}

type Customer struct {
	Id          int64
	UserId      int64
	Email       string
	Name        string
	PhoneNumber string
}

type UserLogin struct {
	Email    string
	Password string
}
