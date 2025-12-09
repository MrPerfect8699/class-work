package ports

import (
	"github.com/yourname/classwork/backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HomeworkService defines the homework service interface
type HomeworkService interface {
	CreateHomework(homework *domain.Homework) error
	GetHomeworksByTeacher(teacherID primitive.ObjectID) ([]domain.Homework, error)
	GetAllHomeworks() ([]domain.Homework, error)
	GetHomeworkByID(id primitive.ObjectID) (*domain.Homework, error)
	UpdateHomework(homework *domain.Homework) error
	DeleteHomework(id primitive.ObjectID) error
}
