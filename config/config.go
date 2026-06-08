package config

import (
	"log"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config holds all application configuration.
type Config struct {
	App      AppConfig
	Postgres PostgresConfig
	MySQL    MySQLConfig
	Mongo    MongoConfig
	Redis    RedisConfig
	JWT      JWTConfig
	Rate     RateConfig
}

type AppConfig struct {
	Env    string `mapstructure:"APP_ENV"`
	Port   string `mapstructure:"APP_PORT"`
	Secret string `mapstructure:"APP_SECRET"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DBName   string `mapstructure:"POSTGRES_DB"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"MYSQL_HOST"`
	Port     string `mapstructure:"MYSQL_PORT"`
	User     string `mapstructure:"MYSQL_USER"`
	Password string `mapstructure:"MYSQL_PASSWORD"`
	DBName   string `mapstructure:"MYSQL_DB"`
}

type MongoConfig struct {
	URI    string `mapstructure:"MONGO_URI"`
	DBName string `mapstructure:"MONGO_DB"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"REDIS_ADDR"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DB"`
}

type JWTConfig struct {
	Secret             string `mapstructure:"JWT_SECRET"`
	ExpiryHours        int    `mapstructure:"JWT_EXPIRY_HOURS"`
	RefreshExpiryHours int    `mapstructure:"JWT_REFRESH_EXPIRY_HOURS"`
}

type RateConfig struct {
	RPS   float64 `mapstructure:"RATE_LIMIT_RPS"`
	Burst int     `mapstructure:"RATE_LIMIT_BURST"`
}

var (
	cfg  *Config
	once sync.Once
)

// Load reads configuration once (singleton pattern).
func Load() *Config {
	once.Do(func() {
		// Load .env file
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using environment variables")
		}

		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		cfg = &Config{}

		cfg.App = AppConfig{
			Env:    viper.GetString("APP_ENV"),
			Port:   viper.GetString("APP_PORT"),
			Secret: viper.GetString("APP_SECRET"),
		}

		cfg.Postgres = PostgresConfig{
			Host:     viper.GetString("POSTGRES_HOST"),
			Port:     viper.GetString("POSTGRES_PORT"),
			User:     viper.GetString("POSTGRES_USER"),
			Password: viper.GetString("POSTGRES_PASSWORD"),
			DBName:   viper.GetString("POSTGRES_DB"),
		}

		cfg.MySQL = MySQLConfig{
			Host:     viper.GetString("MYSQL_HOST"),
			Port:     viper.GetString("MYSQL_PORT"),
			User:     viper.GetString("MYSQL_USER"),
			Password: viper.GetString("MYSQL_PASSWORD"),
			DBName:   viper.GetString("MYSQL_DB"),
		}

		cfg.Mongo = MongoConfig{
			URI:    viper.GetString("MONGO_URI"),
			DBName: viper.GetString("MONGO_DB"),
		}

		cfg.Redis = RedisConfig{
			Addr:     viper.GetString("REDIS_ADDR"),
			Password: viper.GetString("REDIS_PASSWORD"),
			DB:       viper.GetInt("REDIS_DB"),
		}

		cfg.JWT = JWTConfig{
			Secret:             viper.GetString("JWT_SECRET"),
			ExpiryHours:        viper.GetInt("JWT_EXPIRY_HOURS"),
			RefreshExpiryHours: viper.GetInt("JWT_REFRESH_EXPIRY_HOURS"),
		}

		cfg.Rate = RateConfig{
			RPS:   viper.GetFloat64("RATE_LIMIT_RPS"),
			Burst: viper.GetInt("RATE_LIMIT_BURST"),
		}
	})

	return cfg
}