package usecase

import (
	"payment-application/model"
	"payment-application/model/dto"
	"payment-application/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(payload model.UserCredential) error
	FindAllUser(requestPaging dto.PaginationQueryParam) ([]model.UserCredential, dto.PaginationResponse, error)
	FindUser(id string) (model.UserCredential, error)
	UpdateUser(payload model.UserCredential) error
	DeleteUser(id string) error
	FindByUsernameAndPassword(username string, password string) (model.UserCredential, error)
	UpdatePassword(payload model.UserCredential) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func (u *userUsecase) Register(payload model.UserCredential) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	if err != nil {
		return err
	}

	payload.Password = string(bytes)

	checkUser, _ := u.userRepository.FindByUsername(payload.Username)

	if checkUser.Username != "" {
		return fmt.Errorf("username already exist")
	}

	err = u.userRepository.Create(payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) FindAllUser(requestPaging dto.PaginationQueryParam) ([]model.UserCredential, dto.PaginationResponse, error) {
	return u.userRepository.List(requestPaging)
}

func (u *userUsecase) FindUser(id string) (model.UserCredential, error) {
	if id == "" {
		return model.UserCredential{}, fmt.Errorf("id is required")
	}

	user, err := u.userRepository.Get(id)
	if err != nil {
		return model.UserCredential{}, fmt.Errorf("user not found")
	}

	return user, nil
}

func (u *userUsecase) UpdateUser(payload model.UserCredential) error {
	return u.userRepository.Update(payload)
}

func (u *userUsecase) DeleteUser(id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}

	return u.userRepository.Delete(id)
}

func (u *userUsecase) FindByUsernameAndPassword(username string, password string) (model.UserCredential, error) {
	user, err := u.userRepository.FindByUsername(username)
	if err != nil {
		return model.UserCredential{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.UserCredential{}, err
	}

	return user, nil
}

func (u *userUsecase) UpdatePassword(payload model.UserCredential) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	if err != nil {
		return err
	}

	payload.Password = string(bytes)

	return u.userRepository.UpdatePassword(payload)
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}
