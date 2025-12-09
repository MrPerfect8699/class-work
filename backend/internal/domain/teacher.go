package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Teacher represents a teacher entity in the domain
type Teacher struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"` // Never expose password
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updatedAt"`
}

// TeacherRepository defines the interface for teacher persistence
type TeacherRepository interface {
	Save(teacher *Teacher) error
	FindByEmail(email string) (*Teacher, error)
	FindByID(id interface{}) (*Teacher, error)
	Update(teacher *Teacher) error
	Delete(id interface{}) error
}
