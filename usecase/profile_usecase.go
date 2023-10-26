package usecase

import (
	"payment-application/model"
	"payment-application/model/dto"
	"payment-application/repository"
	"errors"
)

type ProfileUsecase interface {
	Create(payload model.Profile) error
	List() ([]model.Profile, error)
	Get(id string) (model.Profile, error)
	Update(payload model.Profile) error
	Delete(id string) error
	Pagination(requestPage dto.PaginationQueryParam, name string) ([]model.Profile, dto.PaginationResponse, error)
}

type profileUsecase struct {
	repository repository.ProfileRepository
}

func (u *profileUsecase) Create(payload model.Profile) error {
	return u.repository.Create(payload)
}

func (u *profileUsecase) List() ([]model.Profile, error) {
	return u.repository.List()
}

func (u *profileUsecase) Get(id string) (model.Profile, error) {
	if id == "" {
		return model.Profile{}, errors.New("id is required")
	}

	profile, err := u.repository.Get(id)
	if err != nil {
		return model.Profile{}, errors.New("profile not found")
	}

	return profile, nil
}

func (u *profileUsecase) Update(payload model.Profile) error {
	return u.repository.Update(payload)
}

func (u *profileUsecase) Delete(id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	return u.repository.Delete(id)
}

func (u *profileUsecase) Pagination(requestPage dto.PaginationQueryParam, name string) ([]model.Profile, dto.PaginationResponse, error) {
	return u.repository.Pagination(requestPage, name)
}

func NewProfileUsecase(repository repository.ProfileRepository) ProfileUsecase {
	return &profileUsecase{repository}
}
