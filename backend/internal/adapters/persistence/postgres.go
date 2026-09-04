package persistence

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/yourname/classwork/backend/config"
)

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

// initSchema creates required tables and indexes if they do not exist
func (p *PostgresDB) initSchema(ctx context.Context) error {
	schema := `
	CREATE TABLE IF NOT EXISTS teachers (
		id SERIAL PRIMARY KEY,
		teacher_id VARCHAR(50) UNIQUE NOT NULL,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		mobile VARCHAR(30) DEFAULT '',
		password VARCHAR(255) NOT NULL,
		department VARCHAR(100) DEFAULT '',
		designation VARCHAR(100) DEFAULT 'Teacher',
		qualification VARCHAR(255) DEFAULT '',
		experience_years INTEGER DEFAULT 0,
		status VARCHAR(30) NOT NULL DEFAULT 'ACTIVE',
		avatar_url TEXT DEFAULT '',
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	-- Add columns idempotently in case table was created with an earlier schema
	ALTER TABLE teachers ADD COLUMN IF NOT EXISTS teacher_id VARCHAR(50);
	ALTER TABLE teachers ADD COLUMN IF NOT EXISTS mobile VARCHAR(30) DEFAULT '';
	ALTER TABLE teachers ADD COLUMN IF NOT EXISTS department VARCHAR(100) DEFAULT '';
	ALTER TABLE teachers ADD COLUMN IF NOT EXISTS designation VARCHAR(100) DEFAULT 'Teacher';
	ALTER TABLE teachers ADD COLUMN IF NOT EXISTS qualification VARCHAR(255) DEFAULT '';
	ALTER TABLE teachers ADD COLUMN IF NOT EXISTS experience_years INTEGER DEFAULT 0;
	ALTER TABLE teachers ADD COLUMN IF NOT EXISTS status VARCHAR(30) DEFAULT 'ACTIVE';
	ALTER TABLE teachers ADD COLUMN IF NOT EXISTS avatar_url TEXT DEFAULT '';

	CREATE INDEX IF NOT EXISTS idx_teachers_teacher_id ON teachers(teacher_id);
	CREATE INDEX IF NOT EXISTS idx_teachers_email ON teachers(email);
	CREATE INDEX IF NOT EXISTS idx_teachers_mobile ON teachers(mobile);

	CREATE TABLE IF NOT EXISTS homework (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT DEFAULT '',
		class_name VARCHAR(255) NOT NULL,
		subject VARCHAR(255) NOT NULL,
		teacher_id INTEGER NOT NULL REFERENCES teachers(id) ON DELETE CASCADE,
		attachments TEXT DEFAULT '',
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS submissions (
		id SERIAL PRIMARY KEY,
		homework_id INTEGER NOT NULL REFERENCES homework(id) ON DELETE CASCADE,
		student_name VARCHAR(255) NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT FALSE,
		notes TEXT DEFAULT '',
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_homework_teacher_id ON homework(teacher_id);
	CREATE INDEX IF NOT EXISTS idx_submissions_homework_id ON submissions(homework_id);
	`

	_, err := p.db.ExecContext(ctx, schema)
	return err
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
