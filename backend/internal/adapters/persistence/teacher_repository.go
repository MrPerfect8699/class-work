package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/yourname/classwork/backend/internal/domain"
)

// TeacherRepositoryImpl implements the domain.TeacherRepository interface using PostgreSQL
type TeacherRepositoryImpl struct {
	db *sql.DB
}

// NewTeacherRepository creates a new instance of TeacherRepositoryImpl
func NewTeacherRepository(db *sql.DB) *TeacherRepositoryImpl {
	return &TeacherRepositoryImpl{
		db: db,
	}
}

// Save inserts a new teacher into the database and populates the generated ID
func (r *TeacherRepositoryImpl) Save(teacher *domain.Teacher) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if teacher.Status == "" {
		teacher.Status = "ACTIVE"
	}
	if teacher.Designation == "" {
		teacher.Designation = "Teacher"
	}

	query := `
		INSERT INTO teachers (
			teacher_id, name, email, mobile, password, department, designation,
			qualification, experience_years, status, avatar_url, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id
	`

	err := r.db.QueryRowContext(ctx, query,
		teacher.TeacherID,
		teacher.Name,
		teacher.Email,
		teacher.Mobile,
		teacher.Password,
		teacher.Department,
		teacher.Designation,
		teacher.Qualification,
		teacher.ExperienceYears,
		teacher.Status,
		teacher.AvatarURL,
		teacher.CreatedAt,
		teacher.UpdatedAt,
	).Scan(&teacher.ID)

	if err != nil {
		return fmt.Errorf("failed to insert teacher: %w", err)
	}

	return nil
}

// FindByEmail finds a teacher by their email address
func (r *TeacherRepositoryImpl) FindByEmail(email string) (*domain.Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, COALESCE(teacher_id, ''), name, email, COALESCE(mobile, ''), password,
		       COALESCE(department, ''), COALESCE(designation, 'Teacher'), COALESCE(qualification, ''),
		       COALESCE(experience_years, 0), COALESCE(status, 'ACTIVE'), COALESCE(avatar_url, ''),
		       created_at, updated_at
		FROM teachers
		WHERE email = $1
	`

	var teacher domain.Teacher
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&teacher.ID,
		&teacher.TeacherID,
		&teacher.Name,
		&teacher.Email,
		&teacher.Mobile,
		&teacher.Password,
		&teacher.Department,
		&teacher.Designation,
		&teacher.Qualification,
		&teacher.ExperienceYears,
		&teacher.Status,
		&teacher.AvatarURL,
		&teacher.CreatedAt,
		&teacher.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("teacher not found")
		}
		return nil, fmt.Errorf("failed to find teacher: %w", err)
	}

	return &teacher, nil
}

// FindByTeacherID finds a teacher by their teacher_id registration code
func (r *TeacherRepositoryImpl) FindByTeacherID(teacherID string) (*domain.Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, COALESCE(teacher_id, ''), name, email, COALESCE(mobile, ''), password,
		       COALESCE(department, ''), COALESCE(designation, 'Teacher'), COALESCE(qualification, ''),
		       COALESCE(experience_years, 0), COALESCE(status, 'ACTIVE'), COALESCE(avatar_url, ''),
		       created_at, updated_at
		FROM teachers
		WHERE teacher_id = $1
	`

	var teacher domain.Teacher
	err := r.db.QueryRowContext(ctx, query, teacherID).Scan(
		&teacher.ID,
		&teacher.TeacherID,
		&teacher.Name,
		&teacher.Email,
		&teacher.Mobile,
		&teacher.Password,
		&teacher.Department,
		&teacher.Designation,
		&teacher.Qualification,
		&teacher.ExperienceYears,
		&teacher.Status,
		&teacher.AvatarURL,
		&teacher.CreatedAt,
		&teacher.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("teacher not found")
		}
		return nil, fmt.Errorf("failed to find teacher: %w", err)
	}

	return &teacher, nil
}

// FindByID finds a teacher by their integer ID
func (r *TeacherRepositoryImpl) FindByID(id int64) (*domain.Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, COALESCE(teacher_id, ''), name, email, COALESCE(mobile, ''), password,
		       COALESCE(department, ''), COALESCE(designation, 'Teacher'), COALESCE(qualification, ''),
		       COALESCE(experience_years, 0), COALESCE(status, 'ACTIVE'), COALESCE(avatar_url, ''),
		       created_at, updated_at
		FROM teachers
		WHERE id = $1
	`

	var teacher domain.Teacher
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&teacher.ID,
		&teacher.TeacherID,
		&teacher.Name,
		&teacher.Email,
		&teacher.Mobile,
		&teacher.Password,
		&teacher.Department,
		&teacher.Designation,
		&teacher.Qualification,
		&teacher.ExperienceYears,
		&teacher.Status,
		&teacher.AvatarURL,
		&teacher.CreatedAt,
		&teacher.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("teacher not found")
		}
		return nil, fmt.Errorf("failed to find teacher: %w", err)
	}

	return &teacher, nil
}

// Update updates an existing teacher's records
func (r *TeacherRepositoryImpl) Update(teacher *domain.Teacher) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE teachers
		SET teacher_id = $1, name = $2, email = $3, mobile = $4, password = $5,
		    department = $6, designation = $7, qualification = $8, experience_years = $9,
		    status = $10, avatar_url = $11, updated_at = $12
		WHERE id = $13
	`

	res, err := r.db.ExecContext(ctx, query,
		teacher.TeacherID,
		teacher.Name,
		teacher.Email,
		teacher.Mobile,
		teacher.Password,
		teacher.Department,
		teacher.Designation,
		teacher.Qualification,
		teacher.ExperienceYears,
		teacher.Status,
		teacher.AvatarURL,
		time.Now(),
		teacher.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update teacher: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("teacher not found")
	}

	return nil
}

// Delete deletes a teacher by ID
func (r *TeacherRepositoryImpl) Delete(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM teachers WHERE id = $1`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete teacher: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("teacher not found")
	}

	return nil
}
