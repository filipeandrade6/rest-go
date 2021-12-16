package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Auth struct {
	Directory string `mapstructure:"directory"`
	ActiveKID string `mapstructure:"activekid"`
}

type Database struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Name         string `mapstructure:"name"`
	MaxIDLEConns int    `mapstructure:"maxidleconns"`
	MaxOpenConns int    `mapstructure:"maxopenconns"`
	DisableTLS   bool   `mapstructure:"disabletls"`
}

type Service struct {
	Host string `mapstructure:"host"`
	Conn string `mapstructure:"conn"`
	Port int    `mapstructure:"port"`
}

type Configuration struct {
	Auth     Auth     `mapstructure:"auth"`
	Database Database `mapstructure:"database"`
	Service  Service  `mapstructure:"service"`
}

func ParseConfig(build string) (Configuration, error) {
	viper.SetDefault("database.host", "dev-postgres:5432")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "secret")
	viper.SetDefault("database.name", "vigia")
	viper.SetDefault("database.maxidleconns", "0")
	viper.SetDefault("database.maxopenconns", "0")
	viper.SetDefault("database.disabletls", "true")

	viper.SetDefault("service.host", "gerencia")
	viper.SetDefault("service.conn", "tcp")
	viper.SetDefault("service.port", "12346")

	viper.BindEnv("database.host", "VIGIA_DB_HOST")
	viper.BindEnv("database.user", "VIGIA_DB_USER")
	viper.BindEnv("database.password", "VIGIA_DB_PASSWORD")
	viper.BindEnv("database.name", "VIGIA_DB_NAME")
	viper.BindEnv("database.maxidleconns", "VIGIA_DB_MAXIDLECONNS")
	viper.BindEnv("database.maxopenconns", "VIGIA_DB_MAXOPENCONNS")
	viper.BindEnv("database.disabletls", "VIGIA_DB_DISABLETLS")

	viper.BindEnv("service.host", "VIGIA_GER_HOST")
	viper.BindEnv("service.conn", "VIGIA_GER_SERVER_CONN")
	viper.BindEnv("service.port", "VIGIA_GER_SERVER_PORT")

	viper.AutomaticEnv()
	fmt.Println(viper.Get("service.host")) // TODO alterar isso aqui

	cfg := Configuration{}

	if err := viper.Unmarshal(&cfg); err != nil {
		return Configuration{}, fmt.Errorf("unmsarshalling config: %w", err)
	}

	return cfg, nil
}

// TODO colocar um metodo pretty print para logar
