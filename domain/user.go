package domain

import (
	"fmt"
	"time"
	"unicode/utf8"
)

// Userのentity
type User struct {
	ID int

	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Validate() error {
	if len(u.Name) == 0 {
		return fmt.Errorf("User name is empty")
	}

	if utf8.RuneCountInString(u.Name) > 20 {
		return fmt.Errorf("User name is too long")
	}

	return nil
}

// Userのrepository
type UserRepository interface {
	Find(id int) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}
