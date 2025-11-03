package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is a thin wrapper around a *gorm.DB providing convenience read/write helpers
// and lifecycle management for a Postgres connection.
type DB struct {
	Conn *gorm.DB
}

// NewPostgres creates a new gorm DB connection to Postgres using the provided DSN
// and connection pool settings.
func NewPostgres(dsn string, maxOpenConns, maxIdleConns int, connMaxLifetime time.Duration, logLevel logger.LogLevel) (*DB, error) {
	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	gdb, err := gorm.Open(postgres.Open(dsn), cfg)
	if err != nil {
		return nil, err
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	return &DB{Conn: gdb}, nil
}

// Close closes the underlying sql.DB connection pool.
func (db *DB) Close() error {
	sqlDB, err := db.Conn.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// Read executes a raw SQL query and scans the result into dest.
// dest should be a pointer to a struct or slice of structs.
func (db *DB) Read(dest interface{}, query string, args ...interface{}) error {
	return db.Conn.Raw(query, args...).Scan(dest).Error
}

// Write inserts the provided model into the database. The model can be a struct or
// slice of structs. It returns an error on failure.
func (db *DB) Write(model interface{}) error {
	return db.Conn.Create(model).Error
}

// Update updates the given model using gorm's Save (useful when model has primary key set).
func (db *DB) Update(model interface{}) error {
	return db.Conn.Save(model).Error
}

// Delete removes the provided model from the database.
func (db *DB) Delete(model interface{}) error {
	return db.Conn.Delete(model).Error
}
