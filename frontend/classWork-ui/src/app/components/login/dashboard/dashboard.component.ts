import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';
import { HomeworkFormComponent } from '../homework-form/homework-form.component';
import { HomeworkListComponent } from '../homework-list/homework-list.component';
import { Homework } from '../../../entities/models';
import { HomeworkService } from '../../../services/homework.service';
import { AuthService } from '../../../core/auth.service';

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
  ],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.scss',
})
export class DashboardComponent implements OnInit {
  homeworks: Homework[] = [];
  isLoading = false;

  constructor(
    private hw: HomeworkService,
    private auth: AuthService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.loadHomeworks();
  }

  loadHomeworks(): void {
    this.isLoading = true;
    this.hw.list().subscribe({
      next: (data) => {
        this.isLoading = false;
        this.homeworks = data || [];
      },
      error: () => {
        this.isLoading = false;
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

