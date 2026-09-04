import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatChipsModule } from '@angular/material/chips';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { Homework } from '../../../entities/models';
import { HomeworkService } from '../../../services/homework.service';

@Component({
  selector: 'app-homework-list',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatCardModule,
    MatButtonModule,
    MatIconModule,
    MatChipsModule,
    MatProgressSpinnerModule,
    MatInputModule,
    MatFormFieldModule,
  ],
  templateUrl: './homework-list.component.html',
  styleUrl: './homework-list.component.scss',
})
export class HomeworkListComponent implements OnInit {
  @Input() homeworks: Homework[] = [];
  @Input() isLoading = false;
  @Output() refreshRequested = new EventEmitter<void>();

  searchTerm = '';
  deletingId: number | null = null;
  errorMessage: string | null = null;

  constructor(private hwService: HomeworkService) {}

  ngOnInit(): void {}

  get filteredHomeworks(): Homework[] {
    if (!this.searchTerm.trim()) {
      return this.homeworks;
    }
    const term = this.searchTerm.toLowerCase();
    return this.homeworks.filter(
      (h) =>
        h.title?.toLowerCase().includes(term) ||
        h.className?.toLowerCase().includes(term) ||
        h.subject?.toLowerCase().includes(term) ||
        h.description?.toLowerCase().includes(term)
    );
  }

  deleteHomework(id?: number): void {
    if (!id) return;
    if (!confirm('Are you sure you want to delete this homework assignment?')) {
      return;
    }

    this.deletingId = id;
    this.errorMessage = null;

    this.hwService.delete(id).subscribe({
      next: () => {
        this.deletingId = null;
        this.refreshRequested.emit();
      },
      error: (err) => {
        this.deletingId = null;
        this.errorMessage =
          err?.error?.error || 'Failed to delete assignment. Please try again.';
      },
    });
  }

  getSubjectColor(subject?: string): string {
    const s = subject?.toLowerCase() || '';
    if (s.includes('math')) return 'subject-math';
    if (s.includes('physic') || s.includes('chem') || s.includes('bio') || s.includes('sci'))
      return 'subject-science';
    if (s.includes('eng') || s.includes('lit')) return 'subject-english';
    if (s.includes('comp') || s.includes('tech')) return 'subject-tech';
    return 'subject-default';
  }
}

