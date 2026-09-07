import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { GenerateAssignmentRequest, GenerateAssignmentResponse } from '../entities/models';
import { environment } from '../../environment';

@Injectable({ providedIn: 'root' })
export class AiService {
  private readonly http = inject(HttpClient);
  readonly base = environment.apiBase;

  generateAssignment(req: GenerateAssignmentRequest): Observable<GenerateAssignmentResponse> {
    return this.http.post<GenerateAssignmentResponse>(`${this.base}/ai/assignments`, req);
  }
}
