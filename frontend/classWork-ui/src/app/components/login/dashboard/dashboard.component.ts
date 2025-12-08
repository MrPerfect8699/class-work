import { Component } from '@angular/core';
import { HomeworkService } from '../../../services/homework.service';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card';
import { HomeworkFormComponent } from '../homework-form/homework-form.component';
import { HomeworkListComponent } from '../homework-list/homework-list.component';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule, MatCardModule, HomeworkFormComponent, HomeworkListComponent],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.scss',
})
export class DashboardComponent {
  teacherId = 1; // in MVP, set manually or parse token
  homeworks: any[] = [];
  constructor(private hw: HomeworkService){}
  ngOnInit(){ this.load(); }
  load(){ this.hw.list(this.teacherId).subscribe((r:any) => this.homeworks = r); }
}
