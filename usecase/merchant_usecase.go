package usecase

import (
	"payment-application/model"
	"payment-application/model/dto"
	"payment-application/repository"
	"errors"
)

type MerchantUsecase interface {
	Create(payload model.Merchant) error
	List() ([]model.Merchant, error)
	Get(id string) (model.Merchant, error)
	Update(payload model.Merchant) error
	Delete(id string) error
	Pagination(requestPage dto.PaginationQueryParam, name string) ([]model.Merchant, dto.PaginationResponse, error)
}

type merchantUsecase struct {
	repository repository.MerchantRepository
}

func (u *merchantUsecase) Create(payload model.Merchant) error {
	return u.repository.Create(payload)
}

func (u *merchantUsecase) List() ([]model.Merchant, error) {
	return u.repository.List()
}

func (u *merchantUsecase) Get(id string) (model.Merchant, error) {
	if id == "" {
		return model.Merchant{}, errors.New("id is required")
	}

	merchant, err := u.repository.Get(id)
	if err != nil {
		return model.Merchant{}, errors.New("merchant not found")
	}

	return merchant, nil
}

func (u *merchantUsecase) Update(payload model.Merchant) error {
	return u.repository.Update(payload)
}

func (u *merchantUsecase) Delete(id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	return u.repository.Delete(id)
}

func (u *merchantUsecase) Pagination(requestPage dto.PaginationQueryParam, name string) ([]model.Merchant, dto.PaginationResponse, error) {
	return u.repository.Pagination(requestPage, name)
}

func NewMerchantUsecase(repository repository.MerchantRepository) MerchantUsecase {
	return &merchantUsecase{repository}
}
