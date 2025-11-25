package database

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config holds database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// NewConfig creates a new database config from environment variables
func NewConfig() *Config {
	return &Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", ""),
		Database: getEnv("DB_NAME", "todo_app"),
	}
}

// Connect establishes a connection to MySQL database
func Connect(cfg *Config) (*gorm.DB, error) {
	if err := runMigrations(cfg); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	dsn := buildDSN(cfg)

	db, err := gorm.Open(mysqlDriver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func runMigrations(cfg *Config) error {
	dsn := buildDSN(cfg)
	fmt.Print(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	absPath, _ := filepath.Abs("internal/database/migrations")
	migrationURL := (&url.URL{
		Scheme: "file",
		Path:   filepath.ToSlash(absPath), // chuyển \ thành /
	}).String()
	m, err := migrate.NewWithDatabaseInstance(
		migrationURL,
		"mysql",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	return nil
}

// buildDSN builds the Data Source Name for MySQL connection
func buildDSN(cfg *Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
