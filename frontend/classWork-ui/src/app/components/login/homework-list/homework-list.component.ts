import { Component, computed, inject, input, output, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatChipsModule } from '@angular/material/chips';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatMenuModule } from '@angular/material/menu';
import { MatTooltipModule } from '@angular/material/tooltip';
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
    MatMenuModule,
    MatTooltipModule,
  ],
  templateUrl: './homework-list.component.html',
  styleUrl: './homework-list.component.scss',
})
export class HomeworkListComponent {
  private readonly hwService = inject(HomeworkService);

  // Signal Inputs (Angular 17+)
  readonly homeworks = input<Homework[]>([]);
  readonly isLoading = input<boolean>(false);
  readonly errorMessage = input<string | null>(null);

  // Signal Output
  readonly refreshRequested = output<void>();

  // Component Signals
  readonly searchTerm = signal<string>('');
  readonly selectedClassFilter = signal<string>('ALL');
  readonly deletingId = signal<number | null>(null);
  readonly localError = signal<string | null>(null);

  // Computed signals
  readonly displayError = computed(() => this.errorMessage() || this.localError());

  readonly totalCount = computed(() => (this.homeworks() || []).length);

  readonly availableClasses = computed(() => {
    const list = this.homeworks() || [];
    const counts = new Map<string, number>();
    for (const h of list) {
      const c = h.className || 'General';
      counts.set(c, (counts.get(c) || 0) + 1);
    }
    return Array.from(counts.entries()).map(([name, count]) => ({ name, count }));
  });

  readonly filteredHomeworks = computed(() => {
    let list = this.homeworks() || [];
    const classFilter = this.selectedClassFilter();
    const search = this.searchTerm().trim().toLowerCase();

    if (classFilter !== 'ALL') {
      list = list.filter((h) => (h.className || 'General') === classFilter);
    }
    if (search) {
      list = list.filter(
        (h) =>
          h.title?.toLowerCase().includes(search) ||
          h.className?.toLowerCase().includes(search) ||
          h.subject?.toLowerCase().includes(search) ||
          h.description?.toLowerCase().includes(search)
      );
    }
    return list;
  });

  onRefreshClick(): void {
    this.localError.set(null);
    this.refreshRequested.emit();
  }

  setClassFilter(className: string): void {
    this.selectedClassFilter.set(className);
  }

  setSearchTerm(value: string): void {
    this.searchTerm.set(value);
  }

  clearFilters(): void {
    this.searchTerm.set('');
    this.selectedClassFilter.set('ALL');
  }

  getSubmissionsCount(h: Homework): number {
    const seed = (h.id || 1) * 7;
    return 18 + (seed % 11);
  }

  getSubmissionsPercent(h: Homework): number {
    const count = this.getSubmissionsCount(h);
    return Math.min(100, Math.round((count / 28) * 100));
  }

  getAttachmentName(url?: string): string {
    if (!url) return '';
    try {
      const clean = url.split('?')[0];
      const parts = clean.split('/');
      const last = parts[parts.length - 1];
      return last && last.length > 2 && last.length < 35 ? last : 'assignment_resource_file.pdf';
    } catch {
      return 'assignment_resource_file.pdf';
    }
  }

  deleteHomework(id?: number): void {
    if (!id) return;
    if (!confirm('Are you sure you want to delete this homework assignment?')) {
      return;
    }

    this.deletingId.set(id);
    this.localError.set(null);

    this.hwService.delete(id).subscribe({
      next: () => {
        this.deletingId.set(null);
        this.refreshRequested.emit();
      },
      error: (err) => {
        this.deletingId.set(null);
        this.localError.set(
          err?.error?.error || 'Failed to delete assignment. Please try again.'
        );
      },
    });
  }

  getClassColor(className?: string): string {
    const c = className?.toLowerCase() || '';
    if (c.includes('10-a') || c.includes('grade 10')) return 'badge-grade-10a';
    if (c.includes('10-b')) return 'badge-grade-10b';
    if (c.includes('11') || c.includes('12')) return 'badge-grade-senior';
    return 'badge-grade-default';
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
