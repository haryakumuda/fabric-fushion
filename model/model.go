package model

import "time"

type User struct {
	Id       int64
	Email    string
	Password string
	RoleId   int64
}

type Customer struct {
	Id          int64
	UserId      int64
	Name        string
	PhoneNumber string
}

type UserLogin struct {
	Email    string
	Password string
}

type Employee struct {
	Id         int64
	UserId     int64
	Name       string
	PositionId int64
}

type Products struct {
	ID         int
	Name       string
	Price      float64
	CategoryID int
}

type Sales struct {
	ID         int
	OrderDate  time.Time
	CustomerId int64
}

type Categories struct {
	ID       int
	Category string
}

type SalesProduct struct {
	ID        int
	SaleID    int
	ProductId int
	Quantity  int
}
