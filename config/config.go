package config

import (
	"payment-application/utils/common"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type DBConfig struct {
	Host, Port, Name, User, Password, Driver string
}

type APIConfig struct {
	APIHost, APIPort string
}

type Config struct {
	APIConfig
	DBConfig
	FileConfig
	TokenConfig
}

type FileConfig struct {
	FilePath  string
	ImagePath string
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     []byte
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

func (c *Config) ReadConfig() error {
	err := common.LoadENV()
	if err != nil {
		return err
	}

	c.DBConfig = DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.APIConfig = APIConfig{
		APIHost: os.Getenv("API_HOST"),
		APIPort: os.Getenv("API_PORT"),
	}

	c.FileConfig = FileConfig{
		FilePath:  os.Getenv("FILE_PATH"),
		ImagePath: os.Getenv("IMAGE_PATH"),
	}

	appTokenExpire, err := strconv.Atoi(os.Getenv("APP_TOKEN_EXPIRE"))
	if err != nil {
		return err
	}

	accessTokenLifeTime := time.Duration(appTokenExpire) * time.Minute

	c.TokenConfig = TokenConfig{
		ApplicationName:     os.Getenv("APP_TOKEN_NAME"),
		JwtSignatureKey:     []byte(os.Getenv("APP_TOKEN_KEY")),
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: accessTokenLifeTime,
	}

	if c.DBConfig.Host == "" || c.DBConfig.Port == "" || c.DBConfig.Name == "" || c.DBConfig.User == "" || c.DBConfig.Password == "" || c.DBConfig.Driver == "" || c.APIConfig.APIHost == "" || c.APIConfig.APIPort == "" {
		return fmt.Errorf("missing required enivronment variables")
	}

	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := config.ReadConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}
