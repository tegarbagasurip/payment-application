package usecase

import (
    "errors"
    "payment-application/model"
    "payment-application/model/dto"
    "payment-application/repository"
    "strconv"
)

type TransferUsecase interface {
    Create(payload model.Transfer) error
    List() ([]model.Transfer, error)
    Get(id string) (model.Transfer, error)
    Update(payload model.Transfer) error
    Delete(id string) error
    Pagination(requestPage dto.PaginationQueryParam, description string) ([]model.Transfer, dto.PaginationResponse, error)
    TransferBalance(transfer model.Transfer) error
}

type transferUsecase struct {
    profileRepository  repository.ProfileRepository
    transferRepository repository.TransferRepository
}

func (u *transferUsecase) Create(payload model.Transfer) error {
    return u.transferRepository.Create(payload)
}

func (u *transferUsecase) List() ([]model.Transfer, error) {
    return u.transferRepository.List()
}

func (u *transferUsecase) Get(id string) (model.Transfer, error) {
    if id == "" {
        return model.Transfer{}, errors.New("id is required")
    }

    transfer, err := u.transferRepository.Get(id)
    if err != nil {
        return model.Transfer{}, errors.New("transfer not found")
    }

    return transfer, nil
}

func (u *transferUsecase) Update(payload model.Transfer) error {
    return u.transferRepository.Update(payload)
}

func (u *transferUsecase) Delete(id string) error {
    if id == "" {
        return errors.New("id is required")
    }

    return u.transferRepository.Delete(id)
}

func (u *transferUsecase) Pagination(requestPage dto.PaginationQueryParam, description string) ([]model.Transfer, dto.PaginationResponse, error) {
    return u.transferRepository.Pagination(requestPage, description)
}

func (u *transferUsecase) TransferBalance(transfer model.Transfer) error {
    // Validasi bahwa profil pengirim memiliki saldo yang cukup
    senderProfile, err := u.profileRepository.Get(transfer.SenderID)
    if err != nil {
        return errors.New("Sender profile not found")
    }

    amount, err := strconv.ParseFloat(transfer.Amount, 64)
    if err != nil {
        return err
    }

    senderBalance, err := strconv.ParseFloat(senderProfile.Balance, 64)
    if err != nil {
        return err
    }

    if senderBalance < amount {
        return errors.New("Insufficient balance")
    }

    // Kurangi saldo profil pengirim
    senderBalance -= amount
    senderProfile.Balance = strconv.FormatFloat(senderBalance, 'f', -1, 64)
    err = u.profileRepository.Update(senderProfile)
    if err != nil {
        return err
    }

    // Tambah saldo merchant penerima
    receiverProfile, err := u.profileRepository.Get(transfer.ReceiverID)
    if err != nil {
        return errors.New("Receiver profile not found")
    }

    receiverBalance, err := strconv.ParseFloat(receiverProfile.Balance, 64)
    if err != nil {
        return err
    }

    receiverBalance += amount
    receiverProfile.Balance = strconv.FormatFloat(receiverBalance, 'f', -1, 64)
    err = u.profileRepository.Update(receiverProfile)
    if err != nil {
        return err
    }

    // Simpan transfer ke database
    err = u.transferRepository.Create(transfer)
    if err != nil {
        return err
    }

    return nil
}

func NewTransferUsecase(profileRepo repository.ProfileRepository, transferRepo repository.TransferRepository) TransferUsecase {
    return &transferUsecase{
        profileRepository:  profileRepo,
        transferRepository: transferRepo,
    }
}
