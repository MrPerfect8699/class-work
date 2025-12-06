import { Component } from '@angular/core';
import { HomeworkService } from '../../homework.service';

@Component({
  selector: 'app-homework-list',
  templateUrl: './homework-list.component.html'
})
export class HomeworkListComponent {
  homeworks: any[] = [];
  teacherId = 1;
  constructor(private hw: HomeworkService) {}
  ngOnInit(){ this.hw.list(this.teacherId).subscribe((r:any)=> this.homeworks = r); }
}
