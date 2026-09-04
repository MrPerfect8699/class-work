package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Env      string
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Port         int
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// DatabaseConfig holds PostgreSQL database configuration
type DatabaseConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Name            string
	SSLMode         string
	URL             string
	ConnTimeout     time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// DSN returns the PostgreSQL connection data source name
func (db *DatabaseConfig) DSN() string {
	if db.URL != "" {
		return db.URL
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s connect_timeout=%d",
		db.Host, db.Port, db.User, db.Password, db.Name, db.SSLMode, int(db.ConnTimeout.Seconds()))
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret         string
	ExpirationTime time.Duration
	RefreshTime    time.Duration
	Algorithm      string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file (optional, won't fail if not found)
	_ = godotenv.Load()

	cfg := &Config{
		Env: os.Getenv("ENV"),
		Server: ServerConfig{
			Port:         getEnvInt("PORT", 8080),
			Host:         getEnvString("HOST", "0.0.0.0"),
			ReadTimeout:  getEnvDuration("READ_TIMEOUT", 10*time.Second),
			WriteTimeout: getEnvDuration("WRITE_TIMEOUT", 10*time.Second),
			IdleTimeout:  getEnvDuration("IDLE_TIMEOUT", 120*time.Second),
		},
		Database: DatabaseConfig{
			Host:            getEnvString("DB_HOST", "localhost"),
			Port:            getEnvInt("DB_PORT", 5432),
			User:            getEnvString("DB_USER", "postgres"),
			Password:        getEnvString("DB_PASSWORD", "postgres"),
			Name:            getEnvString("DB_NAME", getEnvString("DATABASE_NAME", "classwork")),
			SSLMode:         getEnvString("DB_SSLMODE", "disable"),
			URL:             getEnvString("DATABASE_URL", getEnvString("POSTGRES_URI", "")),
			ConnTimeout:     getEnvDuration("DB_CONN_TIMEOUT", 10*time.Second),
			MaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: getEnvDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute),
		},
		JWT: JWTConfig{
			Secret:         getEnvString("JWT_SECRET", "your-secret-key-change-in-production"),
			ExpirationTime: getEnvDuration("JWT_EXPIRATION", 72*time.Hour),
			RefreshTime:    getEnvDuration("JWT_REFRESH_TIME", 24*time.Hour),
			Algorithm:      "HS256",
		},
	}

	if cfg.Env == "" {
		cfg.Env = "development"
	}

	return cfg, nil
}

// Helper functions
func getEnvString(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

// LogConfig logs the current configuration (without sensitive data)
func (c *Config) LogConfig() {
	log.Println("=== Application Configuration ===")
	log.Printf("Environment: %s", c.Env)
	log.Printf("Server: %s:%d", c.Server.Host, c.Server.Port)
	log.Printf("Database: %s (%s:%d)", c.Database.Name, c.Database.Host, c.Database.Port)
	log.Printf("Read Timeout: %s", c.Server.ReadTimeout)
	log.Printf("Write Timeout: %s", c.Server.WriteTimeout)
	log.Printf("JWT Expiration: %s", c.JWT.ExpirationTime)
	log.Println("=================================")
}
