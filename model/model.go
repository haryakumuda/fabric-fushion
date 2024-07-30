package model

import (
	"time"
)

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

type Employee struct {
	Id       int64
	UserId   int64
	Email    string
	Name     string
	Position string
}

type Products struct {
	ID       int
	Name     string
	Price    float64
	Category string
}

type Sales struct {
	ID         int
	OrderDate  time.Time
	CustomerId int64
}
