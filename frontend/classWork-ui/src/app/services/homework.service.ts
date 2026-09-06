import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Homework } from '../entities/models';
import { environment } from '../../environment';

@Injectable({ providedIn: 'root' })
export class HomeworkService {
  private readonly http = inject(HttpClient);
  readonly base = environment.apiBase;

  create(hw: Partial<Homework>): Observable<Homework> {
    return this.http.post<Homework>(`${this.base}/homework`, hw);
  }

  list(teacherId?: number): Observable<Homework[]> {
    const query = teacherId ? `?teacherId=${teacherId}` : '';
    return this.http.get<Homework[]>(`${this.base}/homeworks${query}`);
  }

  get(id: number): Observable<Homework> {
    return this.http.get<Homework>(`${this.base}/homework/${id}`);
  }

  update(id: number, hw: Partial<Homework>): Observable<Homework> {
    return this.http.put<Homework>(`${this.base}/homework/${id}`, hw);
  }

  delete(id: number): Observable<{ message: string }> {
    return this.http.delete<{ message: string }>(`${this.base}/homework/${id}`);
  }
}

