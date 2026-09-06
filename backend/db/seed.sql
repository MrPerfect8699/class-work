-- ==========================================================
-- Classwork Database Mock Data Seed (PostgreSQL)
-- ==========================================================

-- 1. Seed Classrooms
INSERT INTO classrooms (name, grade_level, room_number, academic_year) VALUES
('Grade 10-A', 'Grade 10', 'Room 201', '2026-2027'),
('Grade 10-B', 'Grade 10', 'Room 202', '2026-2027'),
('Grade 11-Science', 'Grade 11', 'Lab 3', '2026-2027'),
('Grade 12-Science', 'Grade 12', 'Lab 1', '2026-2027')
ON CONFLICT (name) DO NOTHING;

-- 2. Seed Default Demo Teacher (email: teacher@classwork.edu, password: password123)
-- bcrypt hash for 'password123': $2a$10$n4qGfVlPekyqEcw52JpQ6Oq5pB33fX4F4k7sN5a6Z6s.gW0sA1yOa
INSERT INTO teachers (
    teacher_id, name, email, mobile, password, department, designation, qualification, experience_years, status
) VALUES (
    'TCH-1001',
    'Dr. Jane Smith',
    'teacher@classwork.edu',
    '+1 (555) 234-5678',
    '$2a$10$n4qGfVlPekyqEcw52JpQ6Oq5pB33fX4F4k7sN5a6Z6s.gW0sA1yOa',
    'Mathematics',
    'Senior Teacher',
    'Ph.D. in Applied Mathematics',
    8,
    'ACTIVE'
) ON CONFLICT (email) DO UPDATE SET
    teacher_id = EXCLUDED.teacher_id,
    name = EXCLUDED.name,
    designation = EXCLUDED.designation,
    department = EXCLUDED.department;

-- Seed Secondary Teacher (email: alan@classwork.edu, password: password123)
INSERT INTO teachers (
    teacher_id, name, email, mobile, password, department, designation, qualification, experience_years, status
) VALUES (
    'TCH-1002',
    'Prof. Alan Turing',
    'alan@classwork.edu',
    '+1 (555) 987-6543',
    '$2a$10$n4qGfVlPekyqEcw52JpQ6Oq5pB33fX4F4k7sN5a6Z6s.gW0sA1yOa',
    'Computer Science',
    'Head of Department (HOD)',
    'M.Tech in Computer Science',
    12,
    'ACTIVE'
) ON CONFLICT (email) DO NOTHING;

-- 3. Seed Students
INSERT INTO students (student_code, name, email, class_name) VALUES
('STU-1001', 'Alex Johnson', 'alex.johnson@student.edu', 'Grade 10-A'),
('STU-1002', 'Maya Patel', 'maya.patel@student.edu', 'Grade 10-A'),
('STU-1003', 'Liam Chen', 'liam.chen@student.edu', 'Grade 10-A'),
('STU-1004', 'Sophia Garcia', 'sophia.garcia@student.edu', 'Grade 10-A'),
('STU-1005', 'Noah Williams', 'noah.williams@student.edu', 'Grade 10-B'),
('STU-1006', 'Emma Davis', 'emma.davis@student.edu', 'Grade 10-B'),
('STU-1007', 'Ethan Miller', 'ethan.miller@student.edu', 'Grade 11-Science'),
('STU-1008', 'Olivia Brown', 'olivia.brown@student.edu', 'Grade 12-Science')
ON CONFLICT (student_code) DO NOTHING;

-- 4. Seed Mock Homework Assignments for Teachers (if not already populated)
DO $$
DECLARE
    tid INTEGER;
BEGIN
    FOR tid IN SELECT id FROM teachers LOOP
        IF NOT EXISTS (SELECT 1 FROM homework WHERE teacher_id = tid) THEN
            -- Assignment 1: Linear Algebra
            INSERT INTO homework (title, description, class_name, subject, teacher_id, attachments, created_at, updated_at)
            VALUES (
                'Linear Algebra Problem Set 4',
                'Complete exercises 4.1 to 4.8 on page 142. Make sure to show all working steps for synthetic division and state domain restrictions clearly.',
                'Grade 10-A',
                'Mathematics',
                tid,
                'https://example.com/algebra_worksheet_ch4.pdf',
                CURRENT_TIMESTAMP - INTERVAL '1 day',
                CURRENT_TIMESTAMP - INTERVAL '1 day'
            );

            -- Assignment 2: Quadratic Equations
            INSERT INTO homework (title, description, class_name, subject, teacher_id, attachments, created_at, updated_at)
            VALUES (
                'Quadratic Equations & Graphing Quiz Prep',
                'Practice factoring quadratic equations and finding the vertex and intercepts. Upload handwritten scratchpad or photo of your graphing notebook before the deadline.',
                'Grade 10-A',
                'Mathematics',
                tid,
                '',
                CURRENT_TIMESTAMP - INTERVAL '2 days',
                CURRENT_TIMESTAMP - INTERVAL '2 days'
            );

            -- Assignment 3: Trigonometry Basics
            INSERT INTO homework (title, description, class_name, subject, teacher_id, attachments, created_at, updated_at)
            VALUES (
                'Trigonometry Basics: Sine & Cosine Ratios',
                'Review right-angle triangles from chapter 3. Answer questions 1 through 10 in the attached worksheet. Calculator precision up to 3 decimal places.',
                'Grade 10-B',
                'Mathematics',
                tid,
                'https://example.com/trig_intro_sheet.pdf',
                CURRENT_TIMESTAMP - INTERVAL '3 days',
                CURRENT_TIMESTAMP - INTERVAL '3 days'
            );

            -- Assignment 4: Polynomial Long Division
            INSERT INTO homework (title, description, class_name, subject, teacher_id, attachments, created_at, updated_at)
            VALUES (
                'Polynomial Long Division Practice',
                'Homework set on polynomial remainder theorem and long division exercises 2.1 to 2.4.',
                'Grade 10-A',
                'Mathematics',
                tid,
                'https://example.com/polynomial_division_guide.pdf',
                CURRENT_TIMESTAMP - INTERVAL '5 days',
                CURRENT_TIMESTAMP - INTERVAL '5 days'
            );

            -- Assignment 5: Physics Lab Report
            INSERT INTO homework (title, description, class_name, subject, teacher_id, attachments, created_at, updated_at)
            VALUES (
                'Newton''s Laws of Motion Lab Report',
                'Submit your complete write-up for Experiment 3: Friction and Inclined Planes including error calculations and velocity graphs.',
                'Grade 11-Science',
                'Physics',
                tid,
                'https://example.com/physics_lab_template.pdf',
                CURRENT_TIMESTAMP - INTERVAL '6 days',
                CURRENT_TIMESTAMP - INTERVAL '6 days'
            );

            -- Assignment 6: Computer Science BST
            INSERT INTO homework (title, description, class_name, subject, teacher_id, attachments, created_at, updated_at)
            VALUES (
                'Data Structures: Binary Search Trees',
                'Implement in-order, pre-order, and post-order traversals in Go or Python. Include time complexity analysis.',
                'Grade 12-Science',
                'Computer Science',
                tid,
                'https://example.com/bst_starter_code.zip',
                CURRENT_TIMESTAMP - INTERVAL '7 days',
                CURRENT_TIMESTAMP - INTERVAL '7 days'
            );
        END IF;
    END LOOP;
END $$;

-- 5. Seed Submissions for Homeworks (if not already seeded)
INSERT INTO submissions (homework_id, student_name, student_code, completed, grade, notes, submitted_at)
SELECT h.id, s.name, s.student_code, true, 'A', 'Submitted on time with complete working.', CURRENT_TIMESTAMP
FROM homework h
CROSS JOIN (
    SELECT name, student_code FROM students LIMIT 4
) s
WHERE NOT EXISTS (
    SELECT 1 FROM submissions sub WHERE sub.homework_id = h.id AND sub.student_code = s.student_code
);
