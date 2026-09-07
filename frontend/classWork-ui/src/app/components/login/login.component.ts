import { Component, OnInit, inject, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSelectModule } from '@angular/material/select';
import { MatTabsModule } from '@angular/material/tabs';
import { AuthService } from '../../core/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatIconModule,
    MatProgressSpinnerModule,
    MatSelectModule,
    MatTabsModule,
  ],
})
export class LoginComponent implements OnInit {
  private readonly fb = inject(FormBuilder);
  private readonly auth = inject(AuthService);
  private readonly router = inject(Router);

  loginForm!: FormGroup;
  registerForm!: FormGroup;

  readonly selectedTabIndex = signal<number>(0);
  readonly hideLoginPassword = signal<boolean>(true);
  readonly hideRegisterPassword = signal<boolean>(true);
  readonly isLoading = signal<boolean>(false);
  readonly errorMessage = signal<string | null>(null);
  readonly successMessage = signal<string | null>(null);

  readonly departments = [
    'Mathematics',
    'Science & Technology',
    'English & Literature',
    'Social Studies & History',
    'Computer Science',
    'Languages',
    'Arts & Music',
    'Physical Education',
    'General',
  ];

  readonly designations = [
    'Teacher',
    'Senior Teacher',
    'Head of Department (HOD)',
    'Assistant Teacher',
    'Guest Lecturer',
  ];

  ngOnInit(): void {
    if (this.auth.isAuthenticated()) {
      this.router.navigate(['/dashboard']);
    }

    this.initForms();
  }

  private initForms(): void {
    this.loginForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
    });

    this.registerForm = this.fb.group({
      name: ['', [Validators.required, Validators.minLength(2)]],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      teacherId: [''],
      mobile: ['', [Validators.pattern('^[0-9+\\-\\s()]*$')]],
      department: ['Mathematics'],
      designation: ['Teacher'],
      qualification: [''],
      experienceYears: [0, [Validators.min(0), Validators.max(50)]],
    });
  }

  switchTab(index: number): void {
    this.selectedTabIndex.set(index);
    this.errorMessage.set(null);
    this.successMessage.set(null);
  }

  dismissAlert(): void {
    this.errorMessage.set(null);
    this.successMessage.set(null);
  }

  onLogin(): void {
    if (this.loginForm.invalid) {
      this.loginForm.markAllAsTouched();
      return;
    }

    this.isLoading.set(true);
    this.errorMessage.set(null);
    this.successMessage.set(null);

    const { email, password } = this.loginForm.value;

    this.auth.login(email, password).subscribe({
      next: (res) => {
        this.isLoading.set(false);
        if (res && res.token) {
          this.router.navigate(['/dashboard']);
        } else {
          this.errorMessage.set('Unexpected response from server. Please try again.');
        }
      },
      error: (err) => {
        this.isLoading.set(false);
        this.errorMessage.set(
          err?.error?.error || 'Invalid email or password. Please try again.'
        );
      },
    });
  }

  onRegister(): void {
    if (this.registerForm.invalid) {
      this.registerForm.markAllAsTouched();
      return;
    }

    this.isLoading.set(true);
    this.errorMessage.set(null);
    this.successMessage.set(null);

    const formValue = this.registerForm.value;
    const payload = {
      name: formValue.name?.trim(),
      email: formValue.email?.trim(),
      password: formValue.password,
      teacherId: formValue.teacherId?.trim() || undefined,
      mobile: formValue.mobile?.trim() || undefined,
      department: formValue.department || undefined,
      designation: formValue.designation || undefined,
      qualification: formValue.qualification?.trim() || undefined,
      experienceYears: formValue.experienceYears ? Number(formValue.experienceYears) : 0,
    };

    this.auth.register(payload).subscribe({
      next: (res) => {
        this.isLoading.set(false);
        const assignedId = res.teacherId ? ` (ID: ${res.teacherId})` : '';
        this.successMessage.set(
          `Registration successful${assignedId}! Please sign in with your credentials.`
        );

        this.loginForm.patchValue({ email: payload.email, password: '' });
        this.registerForm.reset({
          department: 'Mathematics',
          designation: 'Teacher',
          experienceYears: 0,
        });
        this.selectedTabIndex.set(0);
      },
      error: (err) => {
        this.isLoading.set(false);
        this.errorMessage.set(
          err?.error?.error || 'Registration failed. Please check the details and try again.'
        );
      },
    });
  }
}
