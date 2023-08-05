package config

import "os"

func getEnvString(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

type DatabaseConfig struct {
	Username string
	Password string
	Name     string
	Port     string
	IP       string
}

type C struct {
	DB DatabaseConfig
}

func LoadConf() *C {
	c := C{
		DB: DatabaseConfig{
			Username: getEnvString("APP_DB_USERNAME", ""),
			Password: getEnvString("APP_DB_PASSWORD", ""),
			Name:     getEnvString("APP_DB_NAME", ""),
			Port:     getEnvString("APP_DB_PORT", ""),
			IP:       getEnvString("APP_DB_IP", ""),
		},
	}
	return &c
}
