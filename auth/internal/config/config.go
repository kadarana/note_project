package config

import (
	"fmt"
	"os"
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
