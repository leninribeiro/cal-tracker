package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type AppConfig struct {
	DBName     string `env:"DB_NAME"`
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPass     string `env:"DB_PASS"`
	ServerPort string `env:"SERVER_PORT" envDefault:"3000"`
}

var C AppConfig

func init() {
	err := env.Parse(&C)
	if err != nil {
		log.Fatalf("Error loading env vars: %s. Terminating", err)
	}
	C.ServerPort = ":" + C.ServerPort
	log.Println("Successfully loaded env vars")
}
