import { Component, OnInit, computed, inject, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';
import { HomeworkFormComponent } from '../homework-form/homework-form.component';
import { HomeworkListComponent } from '../homework-list/homework-list.component';
import { AiGeneratorComponent } from '../ai-generator/ai-generator.component';
import { Homework } from '../../../entities/models';
import { HomeworkService } from '../../../services/homework.service';
import { AuthService } from '../../../core/auth.service';
import { finalize } from 'rxjs';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [
    CommonModule,
    MatToolbarModule,
    MatButtonModule,
    MatIconModule,
    MatCardModule,
    HomeworkFormComponent,
    HomeworkListComponent,
    AiGeneratorComponent,
  ],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.scss',
})
export class DashboardComponent implements OnInit {
  private readonly hw = inject(HomeworkService);
  private readonly auth = inject(AuthService);
  private readonly router = inject(Router);

  // Component Signals
  readonly activeTab = signal<'manual' | 'ai'>('ai');
  readonly homeworks = signal<Homework[]>([]);
  readonly isLoading = signal<boolean>(false);
  readonly errorMessage = signal<string | null>(null);

  // Computed state
  readonly activeCount = computed(() => this.homeworks().length);

  ngOnInit(): void {
    this.loadHomeworks();
  }

  loadHomeworks(): void {
    this.isLoading.set(true);
    this.errorMessage.set(null);

    this.hw.list().pipe(
      finalize(() => {
        this.isLoading.set(false);
      })
    ).subscribe({
      next: (data) => {
        this.homeworks.set(Array.isArray(data) ? [...data] : []);
      },
      error: (err) => {
        console.error('Failed to load assignments:', err);
        this.homeworks.set([]);
        this.errorMessage.set(
          err?.error?.error || 'Failed to load assignments. Please verify the server is running and try again.'
        );
      },
    });
  }

  onHomeworkCreated(): void {
    this.loadHomeworks();
  }

  logout(): void {
    this.auth.logout();
    this.router.navigate(['/']);
  }
}
