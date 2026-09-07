import { Component, OnInit, inject, output, signal } from '@angular/core';
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
  private readonly fb = inject(FormBuilder);
  private readonly hwService = inject(HomeworkService);

  // Signal Output
  readonly homeworkCreated = output<void>();

  // Component Signals
  readonly isLoading = signal<boolean>(false);
  readonly errorMessage = signal<string | null>(null);
  readonly successMessage = signal<string | null>(null);

  homeworkForm!: FormGroup;

  readonly classes = [
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

  readonly subjects = [
    'Mathematics',
    'Physics',
    'Chemistry',
    'Biology',
    'Computer Science',
    'English',
    'History',
    'General',
  ];

  ngOnInit(): void {
    this.initForm();
  }

  private initForm(): void {
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

    this.isLoading.set(true);
    this.errorMessage.set(null);
    this.successMessage.set(null);

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
        this.isLoading.set(false);
        this.successMessage.set('Assignment published successfully!');
        this.homeworkForm.reset({
          className: 'Grade 10-A',
          subject: 'Mathematics',
        });
        this.homeworkCreated.emit();

        // Auto dismiss success after 4s
        setTimeout(() => {
          this.successMessage.set(null);
        }, 4000);
      },
      error: (err) => {
        this.isLoading.set(false);
        this.errorMessage.set(
          err?.error?.error || 'Failed to create homework. Please try again.'
        );
      },
    });
  }

  dismissAlert(): void {
    this.errorMessage.set(null);
    this.successMessage.set(null);
  }
}
