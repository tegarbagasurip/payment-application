package manager

import (
	"payment-application/usecase"
)

type UseCaseManager interface {
	UserUsecase() usecase.UserUsecase
	AuthUsecase() usecase.AuthUsecase
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

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repo,
	}
}
