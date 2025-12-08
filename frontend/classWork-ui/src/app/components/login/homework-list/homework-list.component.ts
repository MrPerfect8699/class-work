import { Component } from '@angular/core';
import { HomeworkService } from '../../../services/homework.service';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-homework-list',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './homework-list.component.html',
  styleUrl: './homework-list.component.scss',
})
export class HomeworkListComponent {
  homeworks: any[] = [];
  teacherId = 1;
  constructor(private hw: HomeworkService) {}
  ngOnInit(){ this.hw.list(this.teacherId).subscribe((r:any)=> this.homeworks = r); }
}
