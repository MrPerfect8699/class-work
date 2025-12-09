package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Homework represents a homework assignment in the domain
type Homework struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	ClassName   string             `bson:"className" json:"className"`
	Subject     string             `bson:"subject" json:"subject"`
	TeacherID   primitive.ObjectID `bson:"teacherId" json:"teacherId"`
	Attachments string             `bson:"attachments" json:"attachments"`
	CreatedAt   time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updatedAt"`
}

// Submission represents a homework submission
type Submission struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	HomeworkID  primitive.ObjectID `bson:"homeworkId" json:"homeworkId"`
	StudentName string             `bson:"studentName" json:"studentName"`
	Completed   bool               `bson:"completed" json:"completed"`
	Notes       string             `bson:"notes" json:"notes"`
	CreatedAt   time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updatedAt"`
}

// HomeworkRepository defines the interface for homework persistence
type HomeworkRepository interface {
	Save(homework *Homework) error
	FindByID(id primitive.ObjectID) (*Homework, error)
	FindByTeacherID(teacherID primitive.ObjectID) ([]Homework, error)
	FindAll() ([]Homework, error)
	Update(homework *Homework) error
	Delete(id primitive.ObjectID) error
}
