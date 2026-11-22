package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort          string
	OllamaURL        string
	OllamaChatModel  string
	OllamaEmbedModel string
	QdrantURL        string
	RedisAddr        string
	RedisPassword    string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil, fmt.Errorf("failed to load .env: %w", err)
		}
	}

	cfg := &Config{
		AppPort:          getEnv("APP_PORT", "8080"),
		OllamaURL:        getEnv("OLLAMA_URL", "http://localhost:11434"),
		OllamaChatModel:  getEnv("OLLAMA_CHAT_MODEL", "qwen3.5:4b"),
		OllamaEmbedModel: getEnv("OLLAMA_EMBED_MODEL", "embeddinggemma:300m"),
		QdrantURL:        getEnv("QDRANT_URL", "http://localhost:6333"),
		RedisAddr:        getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:    getEnv("REDIS_PASSWORD", ""),
	}

	port, err := strconv.Atoi(cfg.AppPort)

	if err != nil || port < 1 || port > 65535 {
		return nil, fmt.Errorf("invalid APP_PORT: %q", cfg.AppPort)
	}
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
