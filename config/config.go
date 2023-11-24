package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

type Config struct {
	DBConfig
	APPConfig
	RedisConfig
	AuthConfig
}

var cfg Config

type AuthConfig struct {
	PublicKey  string `env:"PUBLIC_KEY,required=true"`
	PrivateKey string `env:"PRIVATE_KEY,required=true"`
}

type RedisConfig struct {
	Address  string `env:"REDIS_ADDRESS,required=true"`
	Password string `env:"REDIS_PASSWORD,required=true"`
	Db       uint32 `env:"REDIS_DB,required=true"`
}

type DBConfig struct {
	Host     string `env:"DB_HOST,required=true"`
	User     string `env:"DB_USER,required=true"`
	Password string `env:"DB_PASSWORD,required=true"`
	DbName   string `env:"DB_DBNAME,required=true"`
	Port     string `env:"DB_PORT,required=true"`
	Sslmode  string `env:"DB_SSLMODE,required=true"`
	Timezone string `env:"DB_TIMEZONE,required=true"`
}

type APPConfig struct {
	Host         string `env:"APP_HOST"`
	Port         string `env:"APP_PORT,required=true"`
	ReadTimeout  uint32 `env:"APP_READ_TIMEOUT,required=true"`
	WriteTimeout uint32 `env:"APP_WRITE_TIMEOUT,required=true"`
	IdleTimeout  uint32 `env:"APP_IDLE_TIMEOUT,required=true"`
}

func init() {
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func GetConfig() Config {
	return cfg
}
