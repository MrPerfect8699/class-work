package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/yourname/classwork/backend/internal/domain"
)

// HomeworkRepositoryImpl implements the domain.HomeworkRepository interface using PostgreSQL
type HomeworkRepositoryImpl struct {
	db *sql.DB
}

// NewHomeworkRepository creates a new instance of HomeworkRepositoryImpl
func NewHomeworkRepository(db *sql.DB) *HomeworkRepositoryImpl {
	return &HomeworkRepositoryImpl{
		db: db,
	}
}

// Save saves a homework assignment to the database and populates the generated ID
func (r *HomeworkRepositoryImpl) Save(homework *domain.Homework) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO homework (title, description, class_name, subject, teacher_id, attachments, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	err := r.db.QueryRowContext(ctx, query,
		homework.Title,
		homework.Description,
		homework.ClassName,
		homework.Subject,
		homework.TeacherID,
		homework.Attachments,
		homework.CreatedAt,
		homework.UpdatedAt,
	).Scan(&homework.ID)

	if err != nil {
		return fmt.Errorf("failed to insert homework: %w", err)
	}

	return nil
}

// FindByID finds a homework assignment by ID
func (r *HomeworkRepositoryImpl) FindByID(id int64) (*domain.Homework, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT 
			h.id, h.title, h.description, h.class_name, h.subject, h.teacher_id, h.attachments,
			COALESCE(sub.sub_count, 0) AS submissions_count,
			COALESCE(stu.stu_count, 0) AS total_students,
			h.created_at, h.updated_at
		FROM homework h
		LEFT JOIN (
			SELECT homework_id, COUNT(*) AS sub_count
			FROM submissions
			GROUP BY homework_id
		) sub ON sub.homework_id = h.id
		LEFT JOIN (
			SELECT class_name, COUNT(*) AS stu_count
			FROM students
			GROUP BY class_name
		) stu ON stu.class_name = h.class_name
		WHERE h.id = $1
	`

	var hw domain.Homework
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&hw.ID,
		&hw.Title,
		&hw.Description,
		&hw.ClassName,
		&hw.Subject,
		&hw.TeacherID,
		&hw.Attachments,
		&hw.SubmissionsCount,
		&hw.TotalStudents,
		&hw.CreatedAt,
		&hw.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("homework not found")
		}
		return nil, fmt.Errorf("failed to find homework: %w", err)
	}

	return &hw, nil
}

// FindByTeacherID finds all homework assignments for a specific teacher
func (r *HomeworkRepositoryImpl) FindByTeacherID(teacherID int64) ([]domain.Homework, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT 
			h.id, h.title, h.description, h.class_name, h.subject, h.teacher_id, h.attachments,
			COALESCE(sub.sub_count, 0) AS submissions_count,
			COALESCE(stu.stu_count, 0) AS total_students,
			h.created_at, h.updated_at
		FROM homework h
		LEFT JOIN (
			SELECT homework_id, COUNT(*) AS sub_count
			FROM submissions
			GROUP BY homework_id
		) sub ON sub.homework_id = h.id
		LEFT JOIN (
			SELECT class_name, COUNT(*) AS stu_count
			FROM students
			GROUP BY class_name
		) stu ON stu.class_name = h.class_name
		WHERE h.teacher_id = $1
		ORDER BY h.created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, teacherID)
	if err != nil {
		return nil, fmt.Errorf("failed to find homeworks: %w", err)
	}
	defer rows.Close()

	var homeworks []domain.Homework
	for rows.Next() {
		var hw domain.Homework
		err := rows.Scan(
			&hw.ID,
			&hw.Title,
			&hw.Description,
			&hw.ClassName,
			&hw.Subject,
			&hw.TeacherID,
			&hw.Attachments,
			&hw.SubmissionsCount,
			&hw.TotalStudents,
			&hw.CreatedAt,
			&hw.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan homework: %w", err)
		}
		homeworks = append(homeworks, hw)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating homeworks: %w", err)
	}

	if homeworks == nil {
		homeworks = []domain.Homework{}
	}

	return homeworks, nil
}

// FindAll finds all homework assignments
func (r *HomeworkRepositoryImpl) FindAll() ([]domain.Homework, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT 
			h.id, h.title, h.description, h.class_name, h.subject, h.teacher_id, h.attachments,
			COALESCE(sub.sub_count, 0) AS submissions_count,
			COALESCE(stu.stu_count, 0) AS total_students,
			h.created_at, h.updated_at
		FROM homework h
		LEFT JOIN (
			SELECT homework_id, COUNT(*) AS sub_count
			FROM submissions
			GROUP BY homework_id
		) sub ON sub.homework_id = h.id
		LEFT JOIN (
			SELECT class_name, COUNT(*) AS stu_count
			FROM students
			GROUP BY class_name
		) stu ON stu.class_name = h.class_name
		ORDER BY h.created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to find homeworks: %w", err)
	}
	defer rows.Close()

	var homeworks []domain.Homework
	for rows.Next() {
		var hw domain.Homework
		err := rows.Scan(
			&hw.ID,
			&hw.Title,
			&hw.Description,
			&hw.ClassName,
			&hw.Subject,
			&hw.TeacherID,
			&hw.Attachments,
			&hw.SubmissionsCount,
			&hw.TotalStudents,
			&hw.CreatedAt,
			&hw.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan homework: %w", err)
		}
		homeworks = append(homeworks, hw)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating homeworks: %w", err)
	}

	if homeworks == nil {
		homeworks = []domain.Homework{}
	}

	return homeworks, nil
}

// Update updates an existing homework assignment
func (r *HomeworkRepositoryImpl) Update(homework *domain.Homework) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE homework
		SET title = $1, description = $2, class_name = $3, subject = $4, attachments = $5, updated_at = $6
		WHERE id = $7
	`

	res, err := r.db.ExecContext(ctx, query,
		homework.Title,
		homework.Description,
		homework.ClassName,
		homework.Subject,
		homework.Attachments,
		time.Now(),
		homework.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update homework: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("homework not found")
	}

	return nil
}

// Delete deletes a homework assignment by ID
func (r *HomeworkRepositoryImpl) Delete(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM homework WHERE id = $1`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete homework: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("homework not found")
	}

	return nil
}
