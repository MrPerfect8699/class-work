package config

import (
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

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	URI         string
	Name        string
	ConnTimeout time.Duration
	MaxPoolSize uint64
	MinPoolSize uint64
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
			URI:         getEnvString("MONGODB_URI", "mongodb://localhost:27017"),
			Name:        getEnvString("DATABASE_NAME", "classwork"),
			ConnTimeout: getEnvDuration("DB_CONN_TIMEOUT", 10*time.Second),
			MaxPoolSize: getEnvUint64("DB_MAX_POOL_SIZE", 100),
			MinPoolSize: getEnvUint64("DB_MIN_POOL_SIZE", 10),
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

func getEnvUint64(key string, defaultValue uint64) uint64 {
	if value := os.Getenv(key); value != "" {
		if uintVal, err := strconv.ParseUint(value, 10, 64); err == nil {
			return uintVal
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
	log.Printf("Database: %s", c.Database.Name)
	log.Printf("Read Timeout: %s", c.Server.ReadTimeout)
	log.Printf("Write Timeout: %s", c.Server.WriteTimeout)
	log.Printf("JWT Expiration: %s", c.JWT.ExpirationTime)
	log.Println("=================================")
}
