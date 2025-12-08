import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
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

  create(hw: Homework) {
    return this.http.post(`${this.base}/homework`, hw, this.headers());
  }
  list(teacherId: number) {
    return this.http.get<Homework[]>(`${this.base}/homeworks?teacherId=${teacherId}`, this.headers());
  }
}
