package application

import (
	"fmt"
	"time"

	"github.com/yourname/classwork/backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HomeworkServiceImpl implements the HomeworkService interface
type HomeworkServiceImpl struct {
	homeworkRepo domain.HomeworkRepository
}

// NewHomeworkService creates a new instance of HomeworkServiceImpl
func NewHomeworkService(homeworkRepo domain.HomeworkRepository) *HomeworkServiceImpl {
	return &HomeworkServiceImpl{
		homeworkRepo: homeworkRepo,
	}
}

// CreateHomework creates a new homework assignment
func (s *HomeworkServiceImpl) CreateHomework(homework *domain.Homework) error {
	if homework.Title == "" {
		return fmt.Errorf("title is required")
	}
	if homework.TeacherID == primitive.NilObjectID {
		return fmt.Errorf("teacher_id is required")
	}

	homework.ID = primitive.NewObjectID()
	homework.CreatedAt = time.Now()
	homework.UpdatedAt = time.Now()

	return s.homeworkRepo.Save(homework)
}

// GetHomeworksByTeacher retrieves all homework assignments for a specific teacher
func (s *HomeworkServiceImpl) GetHomeworksByTeacher(teacherID primitive.ObjectID) ([]domain.Homework, error) {
	if teacherID == primitive.NilObjectID {
		return nil, fmt.Errorf("teacher_id is required")
	}

	homeworks, err := s.homeworkRepo.FindByTeacherID(teacherID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch homeworks: %w", err)
	}

	if homeworks == nil {
		homeworks = []domain.Homework{}
	}

	return homeworks, nil
}

// GetAllHomeworks retrieves all homework assignments
func (s *HomeworkServiceImpl) GetAllHomeworks() ([]domain.Homework, error) {
	homeworks, err := s.homeworkRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch homeworks: %w", err)
	}

	if homeworks == nil {
		homeworks = []domain.Homework{}
	}

	return homeworks, nil
}

// GetHomeworkByID retrieves a specific homework assignment by ID
func (s *HomeworkServiceImpl) GetHomeworkByID(id primitive.ObjectID) (*domain.Homework, error) {
	if id == primitive.NilObjectID {
		return nil, fmt.Errorf("homework_id is required")
	}

	homework, err := s.homeworkRepo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch homework: %w", err)
	}

	return homework, nil
}

// UpdateHomework updates an existing homework assignment
func (s *HomeworkServiceImpl) UpdateHomework(homework *domain.Homework) error {
	if homework.ID == primitive.NilObjectID {
		return fmt.Errorf("homework_id is required")
	}

	homework.UpdatedAt = time.Now()
	return s.homeworkRepo.Update(homework)
}

// DeleteHomework deletes a homework assignment
func (s *HomeworkServiceImpl) DeleteHomework(id primitive.ObjectID) error {
	if id == primitive.NilObjectID {
		return fmt.Errorf("homework_id is required")
	}

	return s.homeworkRepo.Delete(id)
}
