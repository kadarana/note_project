package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port                   string // порт на котором будет запущен сервер
	Host                   string // хост на котором будет запущен сервер
	Timeout                int    // таймаут для операций с сервером в секундах
	DBDSN                  string // строка подключения к базе данных, например, "postgres://user:password@localhost:5432/dbname"
	DBTimeout              int    // таймаут для операций с базой данных в секундах
	JWTSecretKey           string // секретный ключ для JWT токена
	AccessTokenExpiration  int    // срок действия access токена в часах
	RefreshTokenExpiration int    // Срок действия refresh токена в часах
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s: %s", "variable of environment not set", key)
	}
	return value, nil
}

func NewConfig() *Config {
	port, err := getEnv("PORT")
	if err != nil {
		fmt.Println("failed to get PORT from variable of environment, using default port")
	}

	host, err := getEnv("HOST")
	if err != nil {
		fmt.Println("failed to get HOST from variable of environment, using default host")
	}

	timeout := 10
	if envValue, err := getEnv("SERVER_TIMEOUT"); err != nil {
		if parsed, parseErr := strconv.Atoi(envValue); parseErr == nil {
			timeout = parsed
		} else {
			fmt.Printf("failed to get SERVER_TIMEOUT from variable of environment, using %d seconds", timeout)
		}
	}

	dbTimeout := 5
	if envValue, err := getEnv("DB_TIMEOUT"); err != nil {
		if parsed, parseErr := strconv.Atoi(envValue); parseErr == nil {
			timeout = parsed
		} else {
			fmt.Printf("failed to get DB_TIMEOUT from variable of environment, using %d seconds", dbTimeout)
		}
	}

	dbDSN := "Пока не реализовано"

	jwtSecretKey, err := getEnv("JWT_SECRET_KEY")
	if err != nil {
		fmt.Println("failed to get JWT_SECRET_KEY from variable of environment, using default value")
	}

	accessTokenExpiration := 24
	if envValue, err := getEnv("JWT_ACCESS_TOKEN_EXPIRATION"); err == nil {
		if parsed, parseErr := strconv.Atoi(envValue); parseErr == nil {
			accessTokenExpiration = parsed
		}
	}

	refreshTokenExpiration := 168
	if envValue, err := getEnv("REFRESH"); err == nil {
		if parsed, parseErr := strconv.Atoi(envValue); parseErr == nil {
			refreshTokenExpiration = parsed
		}
	}

	return &Config{
		Port:                   port,
		Host:                   host,
		Timeout:                timeout,
		DBDSN:                  dbDSN,
		DBTimeout:              dbTimeout,
		JWTSecretKey:           jwtSecretKey,
		AccessTokenExpiration:  accessTokenExpiration,
		RefreshTokenExpiration: refreshTokenExpiration,
	}

}
