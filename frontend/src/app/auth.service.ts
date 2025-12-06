import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';

@Injectable({ providedIn: 'root' })
export class AuthService {
  base = 'http://localhost:8080/api';
  tokenKey = 'cw_token';
  constructor(private http: HttpClient) {}

  register(name: string, email: string, password: string) {
    return this.http.post(`${this.base}/register`, { name, email, password });
  }
  login(email: string, password: string) {
    return this.http.post<{token: string}>(`${this.base}/login`, { email, password }).pipe(
      map(r => {
        if (r && (r as any).token) {
          localStorage.setItem(this.tokenKey, (r as any).token);
        }
        return r;
      })
    );
  }
  getToken() { return localStorage.getItem(this.tokenKey); }
  logout() { localStorage.removeItem(this.tokenKey); }
}
