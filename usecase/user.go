package usecase

import (
	"go_api_sample/domain"
	"time"
)

type UserUsecase interface {
	FindUser(id int) (*domain.User, error)
	CreateUser(input *CreateUserInput) error
	UpdateUser(input *UpdateUserInput) error
	DeleteUser(id int) error
}

type CreateUserInput struct {
	Name string
}

type UpdateUserInput struct {
	ID   int
	Name string
}

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(
	userRepository domain.UserRepository,
) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) FindUser(id int) (*domain.User, error) {
	user, err := u.userRepository.Find(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) CreateUser(input *CreateUserInput) error {
	now := time.Now()
	user := &domain.User{
		Name:      input.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := user.Validate(); err != nil {
		return err
	}

	if err := u.userRepository.Create(user); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) UpdateUser(input *UpdateUserInput) error {
	user, err := u.userRepository.Find(input.ID)
	if err != nil {
		return err
	}

	user.Name = input.Name
	user.UpdatedAt = time.Now()

	if err := user.Validate(); err != nil {
		return err
	}

	if err := u.userRepository.Update(user); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) DeleteUser(id int) error {
	if err := u.userRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
