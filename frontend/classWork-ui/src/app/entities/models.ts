export interface Teacher {
  id?: number;
  teacherId?: string;
  name: string;
  email: string;
  mobile?: string;
  password?: string;
  department?: string;
  designation?: string;
  qualification?: string;
  experienceYears?: number;
  status?: string;
  avatarUrl?: string;
  createdAt?: string;
  updatedAt?: string;
}

export interface Homework {
  id?: number;
  title: string;
  description?: string;
  className?: string;
  subject?: string;
  attachments?: string;
  teacherId?: number;
  createdAt?: string;
  updatedAt?: string;
}
export interface Classroom {
  id?: number;
  name: string;
  gradeLevel: string;
  roomNumber?: string;
  academicYear?: string;
  createdAt?: string;
}

export interface Student {
  id?: number;
  studentCode: string;
  name: string;
  email: string;
  className: string;
  createdAt?: string;
}

export interface Submission {
  id?: number;
  homeworkId: number;
  studentName: string;
  studentCode?: string;
  completed: boolean;
  grade?: string;
  notes?: string;
  submittedAt?: string;
  createdAt?: string;
  updatedAt?: string;
}

export interface GenerateAssignmentRequest {
  class: string;
  subject: string;
  prompt: string;
}

export interface GeneratedQuestion {
  question: string;
  type: 'short_answer' | 'multiple_choice' | 'essay' | 'true_false' | string;
  marks: number;
}

export interface GenerateAssignmentResponse {
  title: string;
  instructions: string;
  questions: GeneratedQuestion[];
  total_marks: number;
}
