import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Homework } from '../entities/models';
import { AuthService } from '../core/auth.service';
import { environment } from '../../environment';

@Injectable({ providedIn: 'root' })
export class HomeworkService {
  base = environment.apiBase;

  constructor(private http: HttpClient, private auth: AuthService) {}

  private headers() {
    const token = this.auth.getToken();
    return { headers: new HttpHeaders({ Authorization: `Bearer ${token}` }) };
  }

  create(hw: Partial<Homework>): Observable<Homework> {
    return this.http.post<Homework>(`${this.base}/homework`, hw, this.headers());
  }

  list(teacherId?: number): Observable<Homework[]> {
    const query = teacherId ? `?teacherId=${teacherId}` : '';
    return this.http.get<Homework[]>(`${this.base}/homeworks${query}`, this.headers());
  }

  get(id: number): Observable<Homework> {
    return this.http.get<Homework>(`${this.base}/homework/${id}`, this.headers());
  }

  delete(id: number): Observable<{ message: string }> {
    return this.http.delete<{ message: string }>(`${this.base}/homework/${id}`, this.headers());
  }
}

