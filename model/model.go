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

type Sales struct {
	ID         int64
	OrderDate  time.Time
	CustomerId int64
}

type Product struct {
	ID         int64
	CategoryID int64
	Name       string
	Price      float64
	Stock      int64
	Category   string
}

type Categories struct {
	ID       int64
	Category string
}

// SalesProduct represents the sales_products table
type SalesProduct struct {
	ID        int64
	SaleID    int64
	ProductID int64
	Quantity  int
}

// SalesProductDetail represents the detail of products in sales_products
type SalesProductDetail struct {
	ID         int64
	SaleID     int64
	ProductID  int64
	Quantity   int
	Name       string
	Price      float64
	CategoryID int64
	Category   string
}
