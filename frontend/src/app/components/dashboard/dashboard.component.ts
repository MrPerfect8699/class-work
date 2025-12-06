import { Component } from '@angular/core';
import { HomeworkService } from '../../homework.service';
import { Router } from '@angular/router';

@Component({
  templateUrl: './dashboard.component.html'
})
export class DashboardComponent {
  teacherId = 1; // in MVP, set manually or parse token
  homeworks: any[] = [];
  constructor(private hw: HomeworkService){}
  ngOnInit(){ this.load(); }
  load(){ this.hw.list(this.teacherId).subscribe((r:any) => this.homeworks = r); }
}
