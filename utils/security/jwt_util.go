package security

import (
	"errors"
	"fmt"
	"payment-application/config"
	"payment-application/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user model.UserCredential) (string, error) {
	cfg, _ := config.NewConfig()

	now := time.Now().UTC()
	end := now.Add(cfg.AccessTokenLifeTime)

	claims := &TokenMyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.Username,
		Role:     user.Role,
	}

	token := jwt.NewWithClaims(cfg.JwtSigningMethod, claims)
	ss, err := token.SignedString(cfg.JwtSignatureKey)
	if err != nil {
		return "", fmt.Errorf("Failed to create access token : %s", err.Error())
	}
	return ss, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	cfg, _ := config.NewConfig()
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != cfg.JwtSigningMethod {
			return nil, fmt.Errorf("Invalid token signing method")
		}
		return cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Invalid parse token sdf : %s", err.Error())
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != cfg.ApplicationName {
		return nil, fmt.Errorf("Invalid token MapClaims")
	}
	return claims, nil
}

func DeleteAccessToken(userID string) error {
    // Check if the user's access token exists and delete it.
	var accessTokens = make(map[string]bool)

    if _, exists := accessTokens[userID]; exists {
        delete(accessTokens, userID)
        return nil
    }

    // Return an error if the access token doesn't exist.
    return errors.New("Access token not found")
}

