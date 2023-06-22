package config

import (
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

type CurityConfig struct {
	Uri string
}

type VaultConfig struct {
	AdminManagementClientUri string
	MicroServiceClientUri    string
	TokenHeader              string
	Token                    string
}

type ToucanMicroServicesConfig struct {
	Curity CurityConfig
	Vault  VaultConfig
}

func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func GetToucanMicroServicesConfig() *ToucanMicroServicesConfig {
	return &ToucanMicroServicesConfig{
		Curity: CurityConfig{
			Uri: os.Getenv("CURITY_USER_MANAGEMENT_URI"),
		},
		Vault: VaultConfig{
			AdminManagementClientUri: os.Getenv("VAULT_ADMIN_MANAGEMENT_CLIENT_URI"),
			MicroServiceClientUri:    os.Getenv("VAULT_MICRO_SERVICE_CLIENT_URI"),
			TokenHeader:              os.Getenv("VAULT_TOKEN_HEADER"),
			Token:                    os.Getenv("VAULT_TOKEN"),
		},
	}
}
