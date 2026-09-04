import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { environment } from '../../environment';
import { Teacher } from '../entities/models';

export interface AuthResponse {
  id?: number;
  teacherId?: string;
  name?: string;
  email?: string;
  token?: string;
  error?: string;
}

@Injectable({ providedIn: 'root' })
export class AuthService {
  base = environment.apiBase;
  tokenKey = 'cw_token';

  constructor(private http: HttpClient) {}

  register(teacher: Partial<Teacher>): Observable<AuthResponse> {
    return this.http.post<AuthResponse>(`${this.base}/register`, teacher);
  }

  login(email: string, password: string): Observable<AuthResponse> {
    return this.http.post<AuthResponse>(`${this.base}/login`, { email, password }).pipe(
      map((response) => {
        if (response && response.token) {
          localStorage.setItem(this.tokenKey, response.token);
        }
        return response;
      })
    );
  }

  getToken(): string | null {
    return typeof localStorage !== 'undefined' ? localStorage.getItem(this.tokenKey) : null;
  }

  isAuthenticated(): boolean {
    return !!this.getToken();
  }

  logout(): void {
    if (typeof localStorage !== 'undefined') {
      localStorage.removeItem(this.tokenKey);
    }
  }
}

