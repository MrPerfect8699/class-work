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

