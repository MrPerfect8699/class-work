package domain

import (
	"time"
)

// Teacher represents a teacher entity in the domain
type Teacher struct {
	ID              int64     `json:"id"`
	TeacherID       string    `json:"teacherId"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Mobile          string    `json:"mobile"`
	Password        string    `json:"-"` // Never expose password
	Department      string    `json:"department"`
	Designation     string    `json:"designation"`
	Qualification   string    `json:"qualification"`
	ExperienceYears int       `json:"experienceYears"`
	Status          string    `json:"status"`
	AvatarURL       string    `json:"avatarUrl"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// TeacherRepository defines the interface for teacher persistence
type TeacherRepository interface {
	Save(teacher *Teacher) error
	FindByEmail(email string) (*Teacher, error)
	FindByTeacherID(teacherID string) (*Teacher, error)
	FindByID(id int64) (*Teacher, error)
	Update(teacher *Teacher) error
	Delete(id int64) error
}
