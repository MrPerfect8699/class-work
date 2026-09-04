package ports

import (
	"github.com/yourname/classwork/backend/internal/domain"
)

// HomeworkService defines the homework service interface
type HomeworkService interface {
	CreateHomework(homework *domain.Homework) error
	GetHomeworksByTeacher(teacherID int64) ([]domain.Homework, error)
	GetAllHomeworks() ([]domain.Homework, error)
	GetHomeworkByID(id int64) (*domain.Homework, error)
	UpdateHomework(homework *domain.Homework) error
	DeleteHomework(id int64) error
}
