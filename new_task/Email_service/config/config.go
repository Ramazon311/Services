package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment      string
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	LogLevel         string
	RPCPort          string
	SMTPHost         string
	SMTPPort         int
	SMTPUser         string
	SMTPUserPass     string
	EmailFromHeader  string
}

// Load ...
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DB", "newpro"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "coder"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "123"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":7070"))

	c.SMTPHost = cast.ToString(getOrReturnDefault("SMTP_HOST", "smtp.gmail.com"))
	c.SMTPPort = cast.ToInt(getOrReturnDefault("SMTP_PORT", 587))

	c.SMTPUser = cast.ToString(getOrReturnDefault("SMTP_USER", "golang7744@gmail.com"))
	c.SMTPUserPass = cast.ToString(getOrReturnDefault("SMTP_USER_PASSWORD", "g123456789?"))
	c.EmailFromHeader = cast.ToString(getOrReturnDefault("EMAIL_FROM_HEADER", "golang7744@gmail.com"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
