package database

import (
	"database/sql"
	"fmt"
	"net/url"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hoainam183/todo-app/internal/config"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect establishes a connection to MySQL database
func Connect(cfg *config.DBConfig) (*gorm.DB, error) {
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

func runMigrations(cfg *config.DBConfig) error {
	dsn := buildDSN(cfg)
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
func buildDSN(cfg *config.DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
}
