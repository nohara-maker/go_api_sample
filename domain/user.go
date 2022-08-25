package domain

import (
	"time"
)

type User struct {
	ID int

	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	Find(id int) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}
