export interface Teacher {
  id?: number;
  name: string;
  email: string;
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
}
