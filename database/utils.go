package database

import (
	"fmt"
)

// PostgresConfig holds simple Postgres connection parameters useful for building a DSN.
type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string // e.g. disable, require
}

// BuildPostgresDSN builds a URL-style Postgres DSN from PostgresConfig.
// Example: postgres://user:pass@localhost:5432/dbname?sslmode=disable
func BuildPostgresDSN(c PostgresConfig) string {
	ssl := c.SSLMode
	if ssl == "" {
		ssl = "disable"
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.User, c.Password, c.Host, c.Port, c.DBName, ssl)
}

// Ping performs a database-level ping to verify connectivity.
func Ping(db *DB) error {
	sqlDB, err := db.Conn.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

// AutoMigrate runs GORM automigrations for the provided models. Pass model struct types
// (e.g. &User{}, &Order{}) to create/update tables.
func AutoMigrate(db *DB, models ...interface{}) error {
	return db.Conn.AutoMigrate(models...)
}
