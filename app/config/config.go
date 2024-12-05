package config

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Init() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	// 验证必要的环境变量
	requiredEnv := []string{"TG_BOT_ID", "TG_ADMIN_ID"}
	for _, key := range requiredEnv {
		if value := os.Getenv(key); value == "" {
			return errors.New("Required environment variable is not set " + key)
		}
	}

	return nil
}

func GetEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return strings.TrimSpace(value)
	}
	return ""
}

func GetTGBotId() string {
	return GetEnv("TG_BOT_ID")
}

func GetTGAdminId() string {
	return GetEnv("TG_ADMIN_ID")
}
