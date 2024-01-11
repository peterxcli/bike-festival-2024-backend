package bootstrap

import (
	"log"

	"github.com/caarlos0/env/v9"
)

type Env struct {
	DB     DBEnv    `envPrefix:"DB_"`
	Redis  RedisEnv `envPrefix:"REDIS_"`
	Server Server   `envPrefix:"SERVER_"`
	JWT    JWTEnv   `envPrefix:"JWT_"`
	Domain string   `env:"DOMAIN"`
}

func NewEnv() *Env {
	var e Env
	if err := env.ParseWithOptions(&e, env.Options{
		RequiredIfNoDef: true,
		Prefix:          "APP",
	}); err != nil {
		log.Fatal(err)
	}
	return &e
}
