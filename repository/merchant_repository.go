package repository

import (
	"database/sql"
	"payment-application/model"
	"payment-application/model/dto"
	"payment-application/utils/common"
	"payment-application/utils/constant"
)

type MerchantRepository interface {
	BaseRepository[model.Merchant]
	BaseRepositoryPaging[model.Merchant]
}

type merchantRepository struct {
	db *sql.DB
}

func (r *merchantRepository) Create(payload model.Merchant) error {
	_, err := r.db.Exec(constant.MERCHANT_INSERT, payload.Id, payload.NameMerchant, payload.Address, payload.Phone, payload.Balance)
	return err
}

func (r *merchantRepository) List() ([]model.Merchant, error) {
	rows, err := r.db.Query(constant.MERCHANT_LIST)
	if err != nil {
		return nil, err
	}

	var merchants []model.Merchant
	for rows.Next() {
		var merchant model.Merchant
		if err := rows.Scan(&merchant.Id, &merchant.NameMerchant, &merchant.Address, &merchant.Phone, &merchant.Balance); err != nil {
			return nil, err
		}
		merchants = append(merchants, merchant)
	}

	return merchants, nil
}

func (r *merchantRepository) Get(id string) (model.Merchant, error) {
	var merchant model.Merchant
	err := r.db.QueryRow(constant.MERCHANT_GET, id).Scan(&merchant.Id, &merchant.NameMerchant, &merchant.Address, &merchant.Phone, &merchant.Balance)
	return merchant, err
}

func (r *merchantRepository) Update(payload model.Merchant) error {
	_, err := r.db.Exec(constant.MERCHANT_UPDATE, payload.NameMerchant, payload.Address, payload.Phone, payload.Id, payload.Balance)
	return err
}

func (r *merchantRepository) Delete(id string) error {
	_, err := r.db.Exec(constant.MERCHANT_DELETE, id)
	return err
}

func (r *merchantRepository) Pagination(queryPage dto.PaginationQueryParam, query ...string) ([]model.Merchant, dto.PaginationResponse, error) {
	pagination := common.CreatePaginationFromQueryParams(queryPage)

	querySelect := "SELECT id, name_merchant, address, phone, balance FROM merchants"
	if query[0] != "" {
		querySelect += ` WHERE name_merchant ilike '%` + query[0] + `%'`
	}
	querySelect += ` LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(querySelect, pagination.LimitRows, pagination.StartIndex)
	if err != nil {
		return nil, dto.PaginationResponse{}, err
	}

	var merchants []model.Merchant
	for rows.Next() {
		var merchant model.Merchant

		err := rows.Scan(&merchant.Id, &merchant.NameMerchant, &merchant.Address, &merchant.Phone, &merchant.Balance)
		if err != nil {
			return nil, dto.PaginationResponse{}, err
		}

		merchants = append(merchants, merchant)
	}

	var totalRows int
	queryCount := "SELECT COUNT(*) FROM merchants"
	if query[0] != "" {
		queryCount += ` WHERE name_merchant like '%` + query[0] + `%'`
	}

	row := r.db.QueryRow(queryCount)
	err = row.Scan(&totalRows)
	if err != nil {
		return nil, dto.PaginationResponse{}, err
	}
	return merchants, common.CreatePaginationResponse(pagination.CurrentPage, pagination.LimitRows, totalRows), nil
}

func NewMerchantRepository(db *sql.DB) MerchantRepository {
	return &merchantRepository{db}
}
