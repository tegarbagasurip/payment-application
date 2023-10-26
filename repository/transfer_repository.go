package repository

import (
    "database/sql"
    "payment-application/model"
    "payment-application/model/dto"
    "payment-application/utils/common"
    "payment-application/utils/constant"
)

type TransferRepository interface {
    Create(payload model.Transfer) error
    List() ([]model.Transfer, error)
    Get(id string) (model.Transfer, error)
    Update(payload model.Transfer) error
    Delete(id string) error
    Pagination(queryPage dto.PaginationQueryParam, query ...string) ([]model.Transfer, dto.PaginationResponse, error)
}

type transferRepository struct {
    db *sql.DB
}

func (r *transferRepository) Create(payload model.Transfer) error {
    _, err := r.db.Exec(constant.TRANSFER_INSERT, payload.Id, payload.SenderID, payload.ReceiverID, payload.Amount, payload.Description)
    return err
}

func (r *transferRepository) List() ([]model.Transfer, error) {
    rows, err := r.db.Query(constant.TRANSFER_LIST)
    if err != nil {
        return nil, err
    }

    var transfers []model.Transfer
    for rows.Next() {
        var transfer model.Transfer
        if err := rows.Scan(&transfer.Id, &transfer.SenderID, &transfer.ReceiverID, &transfer.Amount, &transfer.Description); err != nil {
            return nil, err
        }
        transfers = append(transfers, transfer)
    }

    return transfers, nil
}

func (r *transferRepository) Get(id string) (model.Transfer, error) {
    var transfer model.Transfer
    err := r.db.QueryRow(constant.TRANSFER_GET, id).Scan(&transfer.Id, &transfer.SenderID, &transfer.ReceiverID, &transfer.Amount, &transfer.Description)
    return transfer, err
}

func (r *transferRepository) Update(payload model.Transfer) error {
    _, err := r.db.Exec(constant.TRANSFER_UPDATE, payload.SenderID, payload.ReceiverID, payload.Amount, payload.Description, payload.Id)
    return err
}

func (r *transferRepository) Delete(id string) error {
    _, err := r.db.Exec(constant.TRANSFER_DELETE, id)
    return err
}

func (r *transferRepository) Pagination(queryPage dto.PaginationQueryParam, query ...string) ([]model.Transfer, dto.PaginationResponse, error) {
    pagination := common.CreatePaginationFromQueryParams(queryPage)

    querySelect := "SELECT id, sender_id, receiver_id, amount, description FROM transfers"
    if query[0] != "" {
        querySelect += ` WHERE description ilike '%` + query[0] + `%'`
    }
    querySelect += ` LIMIT $1 OFFSET $2`

    rows, err := r.db.Query(querySelect, pagination.LimitRows, pagination.StartIndex)
    if err != nil {
        return nil, dto.PaginationResponse{}, err
    }

    var transfers []model.Transfer
    for rows.Next() {
        var transfer model.Transfer

        err := rows.Scan(&transfer.Id, &transfer.SenderID, &transfer.ReceiverID, &transfer.Amount, &transfer.Description)
        if err != nil {
            return nil, dto.PaginationResponse{}, err
        }

        transfers = append(transfers, transfer)
    }

    var totalRows int
    queryCount := "SELECT COUNT(*) FROM transfers"
    if query[0] != "" {
        queryCount += ` WHERE description ilike '%` + query[0] + `%'`
    }

    row := r.db.QueryRow(queryCount)
    err = row.Scan(&totalRows)
    if err != nil {
        return nil, dto.PaginationResponse{}, err
    }
    return transfers, common.CreatePaginationResponse(pagination.CurrentPage, pagination.LimitRows, totalRows), nil
}

func NewTransferRepository(db *sql.DB) TransferRepository {
    return &transferRepository{db}
}
