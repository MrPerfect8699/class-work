import { Component } from '@angular/core';
import { HomeworkService } from '../../../services/homework.service';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-homework-form',
  standalone: true,
  imports: [CommonModule, FormsModule, MatFormFieldModule, MatInputModule, MatButtonModule],
  templateUrl: './homework-form.component.html',
  styleUrls: ['./homework-form.component.scss']
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
