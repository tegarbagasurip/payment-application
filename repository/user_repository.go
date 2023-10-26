package repository

import (
	"database/sql"
	"payment-application/model"
	"payment-application/model/dto"
	"payment-application/utils/common"
	"payment-application/utils/constant"
)

type UserRepository interface {
	Create(payload model.UserCredential) error
	List(requestPaging dto.PaginationQueryParam) ([]model.UserCredential, dto.PaginationResponse, error)
	Get(id string) (model.UserCredential, error)
	Update(payload model.UserCredential) error
	Delete(id string) error
	FindByUsername(username string) (model.UserCredential, error)
	UpdatePassword(payload model.UserCredential) error
}

type userRepository struct {
	db *sql.DB
}

func (r *userRepository) Create(payload model.UserCredential) error {
	_, err := r.db.Exec(constant.USER_INSERT, payload.Id, payload.Username, payload.Password, payload.Role)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) List(requestPaging dto.PaginationQueryParam) ([]model.UserCredential, dto.PaginationResponse, error) {
	queryPaging := common.CreatePaginationFromQueryParams(requestPaging)
	rows, err := r.db.Query(constant.USER_LIST, queryPaging.LimitRows, queryPaging.StartIndex)
	if err != nil {
		return nil, dto.PaginationResponse{}, err
	}

	var users []model.UserCredential
	for rows.Next() {
		var user model.UserCredential
		if err := rows.Scan(&user.Id, &user.Username, &user.IsActive, &user.Role); err != nil {
			return nil, dto.PaginationResponse{}, err
		}
		users = append(users, user)
	}

	var totalRows int
	err = r.db.QueryRow(constant.USER_GET_TOTAL_ROWS).Scan(&totalRows)
	if err != nil {
		return nil, dto.PaginationResponse{}, err
	}

	return users, common.CreatePaginationResponse(queryPaging.CurrentPage, queryPaging.LimitRows, totalRows), nil
}

func (r *userRepository) Get(id string) (model.UserCredential, error) {
	var user model.UserCredential
	err := r.db.QueryRow(constant.USER_GET, id).Scan(&user.Id, &user.Username, &user.IsActive, &user.Role)
	if err != nil {
		return model.UserCredential{}, err
	}

	return user, nil
}

func (r *userRepository) Update(payload model.UserCredential) error {
	_, err := r.db.Exec(constant.USER_UPDATE, payload.Username, payload.Role, payload.IsActive, payload.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(id string) error {
	_, err := r.db.Exec(constant.USER_DELETE, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindByUsername(username string) (model.UserCredential, error) {
	var user model.UserCredential
	err := r.db.QueryRow(constant.USER_GET_BY_USERNAME, username).Scan(&user.Id, &user.Username, &user.Password, &user.IsActive, &user.Role)
	if err != nil {
		return model.UserCredential{}, err
	}

	return user, nil
}

func (r *userRepository) UpdatePassword(payload model.UserCredential) error {
	_, err := r.db.Exec(constant.USER_UPDATE_PASSWORD, payload.Password, payload.Id)
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}
