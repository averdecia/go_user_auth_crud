package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// GlobalConfig general for all project
type GlobalConfig struct {
	MongoDB         Database `json:"mongodb"`
	Elastic         Database `json:"elastic"`
	Port            string   `json:"port"`
	DefaultDatabase string   `json:"default_database"`
	JWT             string   `json:"jwt_key"`
}

// Database general config for all database connection
type Database struct {
	Driver   string      `json:"driver"`
	Host     string      `json:"host"`
	Port     string      `json:"port"`
	User     string      `json:"user"`
	Password string      `json:"password"`
	Database string      `json:"database"`
	Options  interface{} `json:"options"`
}

// Config save all global configurations
var Config GlobalConfig

// LoadConfig loads the configuration from json to struct
func LoadConfig(file string, path string) {
	viper.SetConfigName(file)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %v", err)
		Config = GlobalConfig{}
	}
	// Default values
	viper.SetDefault("port", "8080")

	// Read to struct
	err = viper.Unmarshal(&Config)
	if err != nil {
		Config = GlobalConfig{}
	}
}
