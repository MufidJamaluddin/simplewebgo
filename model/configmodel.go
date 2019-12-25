package model

import (
	"fmt"
)

// ServerConfiguration Model Konfigurasi Server
type ServerConfiguration struct {
	Host string
	Port int
	Env string
}

// IsEnvDev : Apakah Masih Env Development
func(sc *ServerConfiguration) IsEnvDev() bool {
	return (*sc).Env != "prod"
}

// DatabaseConfiguration Model Konfigurasi DB
type DatabaseConfiguration struct {
	Engine	 string
	Host	 string
	Name     string
	User     string
	Password string
	Port	 string
	Timezone string
}

// GetConnectionString Mendapatkan String untuk Koneksi
func (dbc *DatabaseConfiguration) GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", (*dbc).User, (*dbc).Password, (*dbc).Host, (*dbc).Port, (*dbc).Name)
}

// ConfigurationModel Model Konfigurasi Aplikasi
type ConfigurationModel struct {
	Server       ServerConfiguration
	Database     DatabaseConfiguration
}