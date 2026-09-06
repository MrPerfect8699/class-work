-- 1. Teachers Table
CREATE TABLE IF NOT EXISTS teachers (
    id SERIAL PRIMARY KEY,
    teacher_id VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    mobile VARCHAR(30) DEFAULT '',
    password VARCHAR(255) NOT NULL,
    department VARCHAR(100) DEFAULT '',
    designation VARCHAR(100) DEFAULT 'Teacher',
    qualification VARCHAR(255) DEFAULT '',
    experience_years INTEGER DEFAULT 0,
    status VARCHAR(30) NOT NULL DEFAULT 'ACTIVE',
    avatar_url TEXT DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE teachers ADD COLUMN IF NOT EXISTS teacher_id VARCHAR(50);
ALTER TABLE teachers ADD COLUMN IF NOT EXISTS mobile VARCHAR(30) DEFAULT '';
ALTER TABLE teachers ADD COLUMN IF NOT EXISTS department VARCHAR(100) DEFAULT '';
ALTER TABLE teachers ADD COLUMN IF NOT EXISTS designation VARCHAR(100) DEFAULT 'Teacher';
ALTER TABLE teachers ADD COLUMN IF NOT EXISTS qualification VARCHAR(255) DEFAULT '';
ALTER TABLE teachers ADD COLUMN IF NOT EXISTS experience_years INTEGER DEFAULT 0;
ALTER TABLE teachers ADD COLUMN IF NOT EXISTS status VARCHAR(30) DEFAULT 'ACTIVE';
ALTER TABLE teachers ADD COLUMN IF NOT EXISTS avatar_url TEXT DEFAULT '';

CREATE INDEX IF NOT EXISTS idx_teachers_teacher_id ON teachers(teacher_id);
CREATE INDEX IF NOT EXISTS idx_teachers_email ON teachers(email);
CREATE INDEX IF NOT EXISTS idx_teachers_mobile ON teachers(mobile);

-- 2. Classrooms Table
CREATE TABLE IF NOT EXISTS classrooms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    grade_level VARCHAR(50) NOT NULL,
    room_number VARCHAR(50) DEFAULT '',
    academic_year VARCHAR(50) DEFAULT '2026-2027',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 3. Students Table
CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY,
    student_code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    class_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_students_class_name ON students(class_name);

-- 4. Homework Table
CREATE TABLE IF NOT EXISTS homework (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT DEFAULT '',
    class_name VARCHAR(255) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    teacher_id INTEGER NOT NULL REFERENCES teachers(id) ON DELETE CASCADE,
    attachments TEXT DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_homework_teacher_id ON homework(teacher_id);

-- 5. Submissions Table
CREATE TABLE IF NOT EXISTS submissions (
    id SERIAL PRIMARY KEY,
    homework_id INTEGER NOT NULL REFERENCES homework(id) ON DELETE CASCADE,
    student_name VARCHAR(255) NOT NULL,
    student_code VARCHAR(50) DEFAULT '',
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    grade VARCHAR(20) DEFAULT '',
    notes TEXT DEFAULT '',
    submitted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Idempotent column migrations for existing databases
ALTER TABLE homework ADD COLUMN IF NOT EXISTS attachments TEXT DEFAULT '';

ALTER TABLE submissions ADD COLUMN IF NOT EXISTS student_code VARCHAR(50) DEFAULT '';
ALTER TABLE submissions ADD COLUMN IF NOT EXISTS grade VARCHAR(20) DEFAULT '';
ALTER TABLE submissions ADD COLUMN IF NOT EXISTS submitted_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP;

CREATE INDEX IF NOT EXISTS idx_submissions_homework_id ON submissions(homework_id);

