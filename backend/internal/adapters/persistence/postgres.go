package persistence

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/yourname/classwork/backend/config"
)

//go:embed sql/schema.sql
var schemaSQL string

// PostgresDB represents a PostgreSQL connection
type PostgresDB struct {
	db *sql.DB
}

// NewPostgresDB creates a new PostgreSQL database instance and initializes tables
func NewPostgresDB(cfg *config.DatabaseConfig) (*PostgresDB, error) {
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Configure connection pool
	if cfg.MaxOpenConns > 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.MaxIdleConns > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}

	// Verify connection
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}

	pgDB := &PostgresDB{db: db}

	// Initialize database schema
	if err := pgDB.initSchema(ctx); err != nil {
		return nil, fmt.Errorf("failed to initialize schema: %w", err)
	}

	return pgDB, nil
}

// initSchema creates required tables
func (p *PostgresDB) initSchema(ctx context.Context) error {
	// Execute schema DDL
	if _, err := p.db.ExecContext(ctx, schemaSQL); err != nil {
		return fmt.Errorf("failed to execute schema SQL: %w", err)
	}

	return nil
}

// GetDB returns the underlying sql.DB instance
func (p *PostgresDB) GetDB() *sql.DB {
	return p.db
}

// Close closes the database connection
func (p *PostgresDB) Close(ctx context.Context) error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}

