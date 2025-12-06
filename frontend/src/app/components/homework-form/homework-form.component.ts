import { Component } from '@angular/core';
import { HomeworkService } from '../../homework.service';

@Component({
  selector: 'app-homework-form',
  templateUrl: './homework-form.component.html'
})
export class HomeworkFormComponent {
  title=''; description=''; className=''; subject='';
  teacherId = 1;
  constructor(private hw: HomeworkService) {}
  create(){
    this.hw.create({ title:this.title, description:this.description, className:this.className, subject:this.subject, teacherId:this.teacherId }).subscribe(()=>{
      alert('created'); window.location.reload();
    },()=> alert('err'));
  }
}
