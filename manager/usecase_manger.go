package manager

import (
    "payment-application/usecase"
)

type UseCaseManager interface {
    UserUsecase() usecase.UserUsecase
    AuthUsecase() usecase.AuthUsecase
    ProfileUsecase() usecase.ProfileUsecase
    MerchantUsecase() usecase.MerchantUsecase
    TransferUsecase() usecase.TransferUsecase
}

type useCaseManager struct {
    repoManager RepoManager
}

func (u *useCaseManager) UserUsecase() usecase.UserUsecase {
    return usecase.NewUserUsecase(u.repoManager.UserRepo())
}

func (u *useCaseManager) AuthUsecase() usecase.AuthUsecase {
    return usecase.NewAuthUsecase(u.UserUsecase())
}

func (u *useCaseManager) ProfileUsecase() usecase.ProfileUsecase {
    return usecase.NewProfileUsecase(u.repoManager.ProfileRepo())
}

func (u *useCaseManager) MerchantUsecase() usecase.MerchantUsecase {
    return usecase.NewMerchantUsecase(u.repoManager.MerchantRepo())
}

func (u *useCaseManager) TransferUsecase() usecase.TransferUsecase {
    return usecase.NewTransferUsecase(u.repoManager.ProfileRepo(), u.repoManager.TransferRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
    return &useCaseManager{
        repoManager: repo,
    }
}
