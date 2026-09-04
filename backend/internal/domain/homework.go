package domain

import (
	"time"
)

// Homework represents a homework assignment in the domain
type Homework struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ClassName   string    `json:"className"`
	Subject     string    `json:"subject"`
	TeacherID   int64     `json:"teacherId"`
	Attachments string    `json:"attachments"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Submission represents a homework submission
type Submission struct {
	ID          int64     `json:"id"`
	HomeworkID  int64     `json:"homeworkId"`
	StudentName string    `json:"studentName"`
	Completed   bool      `json:"completed"`
	Notes       string    `json:"notes"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// HomeworkRepository defines the interface for homework persistence
type HomeworkRepository interface {
	Save(homework *Homework) error
	FindByID(id int64) (*Homework, error)
	FindByTeacherID(teacherID int64) ([]Homework, error)
	FindAll() ([]Homework, error)
	Update(homework *Homework) error
	Delete(id int64) error
}
