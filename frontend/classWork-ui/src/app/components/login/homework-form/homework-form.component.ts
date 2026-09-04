import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatSelectModule } from '@angular/material/select';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { HomeworkService } from '../../../services/homework.service';

@Component({
  selector: 'app-homework-form',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatIconModule,
    MatSelectModule,
    MatProgressSpinnerModule,
  ],
  templateUrl: './homework-form.component.html',
  styleUrls: ['./homework-form.component.scss'],
})
export class HomeworkFormComponent implements OnInit {
  @Output() homeworkCreated = new EventEmitter<void>();

  homeworkForm!: FormGroup;
  isLoading = false;
  errorMessage: string | null = null;
  successMessage: string | null = null;

  classes = [
    'Grade 6-A',
    'Grade 7-A',
    'Grade 8-A',
    'Grade 9-A',
    'Grade 10-A',
    'Grade 10-B',
    'Grade 11-Science',
    'Grade 12-Science',
    'Other',
  ];

  subjects = [
    'Mathematics',
    'Physics',
    'Chemistry',
    'Biology',
    'Computer Science',
    'English',
    'History',
    'General',
  ];

  constructor(
    private fb: FormBuilder,
    private hwService: HomeworkService
  ) {}

  ngOnInit(): void {
    this.homeworkForm = this.fb.group({
      title: ['', [Validators.required, Validators.minLength(3)]],
      className: ['Grade 10-A', [Validators.required]],
      subject: ['Mathematics', [Validators.required]],
      description: ['', [Validators.required]],
      attachments: [''],
    });
  }

  onSubmit(): void {
    if (this.homeworkForm.invalid) {
      this.homeworkForm.markAllAsTouched();
      return;
    }

    this.isLoading = true;
    this.errorMessage = null;
    this.successMessage = null;

    const formVal = this.homeworkForm.value;
    const payload = {
      title: formVal.title.trim(),
      className: formVal.className.trim(),
      subject: formVal.subject.trim(),
      description: formVal.description.trim(),
      attachments: formVal.attachments?.trim() || '',
    };

    this.hwService.create(payload).subscribe({
      next: () => {
        this.isLoading = false;
        this.successMessage = 'Assignment published successfully!';
        this.homeworkForm.reset({
          className: 'Grade 10-A',
          subject: 'Mathematics',
        });
        this.homeworkCreated.emit();

        // Auto dismiss success after 4s
        setTimeout(() => {
          this.successMessage = null;
        }, 4000);
      },
      error: (err) => {
        this.isLoading = false;
        this.errorMessage =
          err?.error?.error || 'Failed to create homework. Please try again.';
      },
    });
  }

  dismissAlert(): void {
    this.errorMessage = null;
    this.successMessage = null;
  }
}

