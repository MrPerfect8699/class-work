package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/yourname/classwork/backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TeacherRepositoryImpl implements the TeacherRepository interface
type TeacherRepositoryImpl struct {
	db *mongo.Database
}

// NewTeacherRepository creates a new instance of TeacherRepositoryImpl
func NewTeacherRepository(db *mongo.Database) *TeacherRepositoryImpl {
	return &TeacherRepositoryImpl{
		db: db,
	}
}

// Save saves a teacher to the database
func (r *TeacherRepositoryImpl) Save(teacher *domain.Teacher) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.Collection("teachers").InsertOne(ctx, teacher)
	if err != nil {
		return fmt.Errorf("failed to insert teacher: %w", err)
	}

	return nil
}

// FindByEmail finds a teacher by email
func (r *TeacherRepositoryImpl) FindByEmail(email string) (*domain.Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var teacher domain.Teacher
	err := r.db.Collection("teachers").FindOne(ctx, bson.M{"email": email}).Decode(&teacher)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("teacher not found")
		}
		return nil, fmt.Errorf("failed to find teacher: %w", err)
	}

	return &teacher, nil
}

// FindByID finds a teacher by ID
func (r *TeacherRepositoryImpl) FindByID(id interface{}) (*domain.Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, ok := id.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("invalid teacher ID type")
	}

	var teacher domain.Teacher
	err := r.db.Collection("teachers").FindOne(ctx, bson.M{"_id": objID}).Decode(&teacher)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("teacher not found")
		}
		return nil, fmt.Errorf("failed to find teacher: %w", err)
	}

	return &teacher, nil
}

// Update updates a teacher
func (r *TeacherRepositoryImpl) Update(teacher *domain.Teacher) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.Collection("teachers").UpdateOne(ctx, bson.M{"_id": teacher.ID}, bson.M{
		"$set": teacher,
	})
	if err != nil {
		return fmt.Errorf("failed to update teacher: %w", err)
	}

	return nil
}

// Delete deletes a teacher
func (r *TeacherRepositoryImpl) Delete(id interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.Collection("teachers").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("failed to delete teacher: %w", err)
	}

	return nil
}
