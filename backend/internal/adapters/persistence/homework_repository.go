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

// HomeworkRepositoryImpl implements the HomeworkRepository interface
type HomeworkRepositoryImpl struct {
	db *mongo.Database
}

// NewHomeworkRepository creates a new instance of HomeworkRepositoryImpl
func NewHomeworkRepository(db *mongo.Database) *HomeworkRepositoryImpl {
	return &HomeworkRepositoryImpl{
		db: db,
	}
}

// Save saves a homework to the database
func (r *HomeworkRepositoryImpl) Save(homework *domain.Homework) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.Collection("homework").InsertOne(ctx, homework)
	if err != nil {
		return fmt.Errorf("failed to insert homework: %w", err)
	}

	return nil
}

// FindByID finds a homework by ID
func (r *HomeworkRepositoryImpl) FindByID(id primitive.ObjectID) (*domain.Homework, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var homework domain.Homework
	err := r.db.Collection("homework").FindOne(ctx, bson.M{"_id": id}).Decode(&homework)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("homework not found")
		}
		return nil, fmt.Errorf("failed to find homework: %w", err)
	}

	return &homework, nil
}

// FindByTeacherID finds all homework assignments for a teacher
func (r *HomeworkRepositoryImpl) FindByTeacherID(teacherID primitive.ObjectID) ([]domain.Homework, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.db.Collection("homework").Find(ctx, bson.M{"teacherId": teacherID})
	if err != nil {
		return nil, fmt.Errorf("failed to find homeworks: %w", err)
	}
	defer cursor.Close(ctx)

	var homeworks []domain.Homework
	if err = cursor.All(ctx, &homeworks); err != nil {
		return nil, fmt.Errorf("failed to decode homeworks: %w", err)
	}

	return homeworks, nil
}

// FindAll finds all homework assignments
func (r *HomeworkRepositoryImpl) FindAll() ([]domain.Homework, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.db.Collection("homework").Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find homeworks: %w", err)
	}
	defer cursor.Close(ctx)

	var homeworks []domain.Homework
	if err = cursor.All(ctx, &homeworks); err != nil {
		return nil, fmt.Errorf("failed to decode homeworks: %w", err)
	}

	return homeworks, nil
}

// Update updates a homework
func (r *HomeworkRepositoryImpl) Update(homework *domain.Homework) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.Collection("homework").UpdateOne(ctx, bson.M{"_id": homework.ID}, bson.M{
		"$set": homework,
	})
	if err != nil {
		return fmt.Errorf("failed to update homework: %w", err)
	}

	return nil
}

// Delete deletes a homework
func (r *HomeworkRepositoryImpl) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.Collection("homework").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("failed to delete homework: %w", err)
	}

	return nil
}
