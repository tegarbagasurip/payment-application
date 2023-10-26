package repository

import (
	"database/sql"
	"payment-application/model"
	"payment-application/model/dto"
	"payment-application/utils/common"
	"payment-application/utils/constant"
)

type ProfileRepository interface {
	BaseRepository[model.Profile]
	BaseRepositoryPaging[model.Profile]
}

type profileRepository struct {
	db *sql.DB
}

func (r *profileRepository) Create(payload model.Profile) error {
	_, err := r.db.Exec(constant.PROFILE_INSERT, payload.Id, payload.Name, payload.Address, payload.Phone, payload.Balance)
	return err
}

func (r *profileRepository) List() ([]model.Profile, error) {
	rows, err := r.db.Query(constant.PROFILE_LIST)
	if err != nil {
		return nil, err
	}

	var profiles []model.Profile
	for rows.Next() {
		var profile model.Profile
		if err := rows.Scan(&profile.Id, &profile.Name, &profile.Address, &profile.Phone, &profile.Balance); err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func (r *profileRepository) Get(id string) (model.Profile, error) {
	var profile model.Profile
	err := r.db.QueryRow(constant.PROFILE_GET, id).Scan(&profile.Id, &profile.Name, &profile.Address, &profile.Phone, &profile.Balance)
	return profile, err
}

func (r *profileRepository) Update(payload model.Profile) error {
	_, err := r.db.Exec(constant.PROFILE_UPDATE, payload.Name, payload.Address, payload.Phone, payload.Id, payload.Balance)
	return err
}

func (r *profileRepository) Delete(id string) error {
	_, err := r.db.Exec(constant.PROFILE_DELETE, id)
	return err
}

func (r *profileRepository) Pagination(queryPage dto.PaginationQueryParam, query ...string) ([]model.Profile, dto.PaginationResponse, error) {
	pagination := common.CreatePaginationFromQueryParams(queryPage)

	querySelect := "SELECT id, name, address, phone, balance FROM profiles"
	if query[0] != "" {
		querySelect += ` WHERE name ilike '%` + query[0] + `%'`
	}
	querySelect += ` LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(querySelect, pagination.LimitRows, pagination.StartIndex)
	if err != nil {
		return nil, dto.PaginationResponse{}, err
	}

	var profiles []model.Profile
	for rows.Next() {
		var profile model.Profile

		err := rows.Scan(&profile.Id, &profile.Name, &profile.Address, &profile.Phone, &profile.Balance)
		if err != nil {
			return nil, dto.PaginationResponse{}, err
		}

		profiles = append(profiles, profile)
	}

	var totalRows int
	queryCount := "SELECT COUNT(*) FROM profiles"
	if query[0] != "" {
		queryCount += ` WHERE name ilike '%` + query[0] + `%'`
	}

	row := r.db.QueryRow(queryCount)
	err = row.Scan(&totalRows)
	if err != nil {
		return nil, dto.PaginationResponse{}, err
	}
	return profiles, common.CreatePaginationResponse(pagination.CurrentPage, pagination.LimitRows, totalRows), nil
}


func NewProfileRepository(db *sql.DB) ProfileRepository {
	return &profileRepository{db}
}
