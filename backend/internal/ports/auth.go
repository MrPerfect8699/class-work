package ports

import "go.mongodb.org/mongo-driver/bson/primitive"

// AuthService defines the authentication service interface
type AuthService interface {
	Register(name, email, password string) (primitive.ObjectID, error)
	Login(email, password string) (string, error)
	ValidateToken(token string) (primitive.ObjectID, error)
}
