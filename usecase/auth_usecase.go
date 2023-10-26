package usecase

import (
	"fmt"

	"payment-application/utils/security"
)

type AuthUsecase interface {
	Login(username string, password string) (string, error)
	Logout(token string) error
}

type authUsecase struct {
	userUsecase UserUsecase
}

func (a *authUsecase) Login(username string, password string) (string, error) {
	user, err := a.userUsecase.FindByUsernameAndPassword(username, password)
	if err != nil {
		return "", fmt.Errorf("username or password is wrong or user not active")
	}

	token, err := security.CreateAccessToken(user)
	if err != nil {
		return "", fmt.Errorf("error creating access token: %w", err)
	}

	return token, nil
}

func (a *authUsecase) Logout(token string) error {
    if err := security.DeleteAccessToken(token); err != nil {
        return err
    }

    return nil // Logout berhasil
}

func NewAuthUsecase(userUsecase UserUsecase) AuthUsecase {
	return &authUsecase{
		userUsecase: userUsecase,
	}
}
