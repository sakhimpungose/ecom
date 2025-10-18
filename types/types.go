package types

import "time"

type RegisterUserDto struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

type User struct {
	ID           int    `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserStore interface {
	GetUserByEmailAddress(emailAddress string) (*User, error)
	//GetUserByID(id int) (*User, error)
	CreateUser(User) error
}