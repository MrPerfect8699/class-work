import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Homework } from './models';
import { AuthService } from './auth.service';

@Injectable({ providedIn: 'root' })
export class HomeworkService {
  base = 'http://localhost:8080/api';
  constructor(private http: HttpClient, private auth: AuthService) {}

  private headers() {
    const token = this.auth.getToken();
    return { headers: new HttpHeaders({ Authorization: `Bearer ${token}` }) };
  }

  create(hw: Homework) {
    return this.http.post(`${this.base}/homework`, hw, this.headers());
  }
  list(teacherId: number) {
    return this.http.get<Homework[]>(`${this.base}/homeworks?teacherId=${teacherId}`, this.headers());
  }
}
