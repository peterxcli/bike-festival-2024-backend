package bootstrap

import (
	"log"

	"github.com/caarlos0/env/v9"
)

type envType string

var (
	DevelopmentEnv envType = "development"
	ProductionEnv  envType = "production"
	StageEnv       envType = "stage"
)

type Env struct {
	DB     DBEnv    `envPrefix:"DB_"`
	Redis  RedisEnv `envPrefix:"REDIS_"`
	Server Server   `envPrefix:"SERVER_"`
	JWT    JWTEnv   `envPrefix:"JWT_"`
	Line   LineEnv  `envPrefix:"LINE_"`
	Domain string   `env:"DOMAIN"`
	Env    envType  `env:"ENV" envDefault:"development"`
}

func NewEnv() *Env {
	var e Env
	if err := env.ParseWithOptions(&e, env.Options{
		RequiredIfNoDef: true,
		Prefix:          "APP_",
	}); err != nil {
		log.Fatal(err)
	}
	return &e
}
