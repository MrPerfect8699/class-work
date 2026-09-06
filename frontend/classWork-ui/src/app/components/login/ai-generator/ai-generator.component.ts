import { Component, inject, output, signal, computed } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatSelectModule } from '@angular/material/select';
import { MatChipsModule } from '@angular/material/chips';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatDividerModule } from '@angular/material/divider';
import { AiService } from '../../../services/ai.service';
import { HomeworkService } from '../../../services/homework.service';
import { GeneratedQuestion, GenerateAssignmentResponse } from '../../../entities/models';

type GeneratorStage = 'prompt' | 'generating' | 'review' | 'publishing';

@Component({
  selector: 'app-ai-generator',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatIconModule,
    MatSelectModule,
    MatChipsModule,
    MatProgressSpinnerModule,
    MatTooltipModule,
    MatDividerModule,
  ],
  templateUrl: './ai-generator.component.html',
  styleUrls: ['./ai-generator.component.scss'],
})
export class AiGeneratorComponent {
  private readonly fb = inject(FormBuilder);
  private readonly aiService = inject(AiService);
  private readonly hwService = inject(HomeworkService);

  // Outputs
  readonly assignmentPublished = output<void>();

  // State Signals
  readonly stage = signal<GeneratorStage>('prompt');
  readonly errorMessage = signal<string | null>(null);
  readonly successMessage = signal<string | null>(null);

  // Review State Signals
  readonly reviewTitle = signal<string>('');
  readonly reviewInstructions = signal<string>('');
  readonly reviewQuestions = signal<GeneratedQuestion[]>([]);
  readonly reviewClassName = signal<string>('Grade 10-A');
  readonly reviewSubject = signal<string>('Mathematics');

  // Computed Total Marks
  readonly calculatedTotalMarks = computed(() => {
    return this.reviewQuestions().reduce((sum, q) => sum + (Number(q.marks) || 0), 0);
  });

  promptForm: FormGroup = this.fb.group({
    className: ['Grade 10-A', [Validators.required]],
    subject: ['Mathematics', [Validators.required]],
    prompt: ['', [Validators.required, Validators.minLength(5)]],
  });

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

  readonly questionTypes = [
    { value: 'short_answer', label: 'Short Answer' },
    { value: 'multiple_choice', label: 'Multiple Choice (MCQ)' },
    { value: 'essay', label: 'Essay / Descriptive' },
    { value: 'true_false', label: 'True / False' },
  ];

  readonly promptPresets = [
    {
      label: '📐 Quadratic Equations',
      subject: 'Mathematics',
      text: 'Create 5 questions on quadratic equations with varying difficulty and marks.',
    },
    {
      label: '⚡ Newton\'s Laws of Motion',
      subject: 'Physics',
      text: 'Create 5 questions on Newton\'s Laws of Motion including MCQ, short answer, and calculation questions.',
    },
    {
      label: '🧪 Chemical Bonding',
      subject: 'Chemistry',
      text: 'Generate 4 questions on covalent and ionic bonding with clear marking criteria.',
    },
    {
      label: '🌱 Photosynthesis Quiz',
      subject: 'Biology',
      text: 'Create a 6-question quiz on light-dependent and dark reactions with multiple choice and short answer questions.',
    },
  ];

  applyPreset(preset: { label: string; subject: string; text: string }): void {
    this.promptForm.patchValue({
      subject: preset.subject,
      prompt: preset.text,
    });
  }

  onGenerate(): void {
    if (this.promptForm.invalid) {
      this.promptForm.markAllAsTouched();
      return;
    }

    const { className, subject, prompt } = this.promptForm.value;
    this.reviewClassName.set(className);
    this.reviewSubject.set(subject);

    this.stage.set('generating');
    this.errorMessage.set(null);
    this.successMessage.set(null);

    this.aiService.generateAssignment({ class: className, subject, prompt }).subscribe({
      next: (resp: GenerateAssignmentResponse) => {
        this.reviewTitle.set(resp.title || `${subject} - Practice Assignment`);
        this.reviewInstructions.set(resp.instructions || 'Answer all questions carefully.');
        this.reviewQuestions.set(
          (resp.questions || []).map((q) => ({
            question: q.question,
            type: q.type || 'short_answer',
            marks: Number(q.marks) || 1,
          }))
        );
        this.stage.set('review');
      },
      error: (err) => {
        this.stage.set('prompt');
        this.errorMessage.set(
          err?.error?.error || 'Failed to generate assignment with AI. Please verify your prompt or API key.'
        );
      },
    });
  }

  // --- Teacher Review & Editing Methods ---

  updateQuestionText(index: number, text: string): void {
    const list = [...this.reviewQuestions()];
    if (list[index]) {
      list[index] = { ...list[index], question: text };
      this.reviewQuestions.set(list);
    }
  }

  updateQuestionType(index: number, type: string): void {
    const list = [...this.reviewQuestions()];
    if (list[index]) {
      list[index] = { ...list[index], type };
      this.reviewQuestions.set(list);
    }
  }

  updateQuestionMarks(index: number, marks: number): void {
    const validMarks = Math.max(1, Math.min(100, Number(marks) || 1));
    const list = [...this.reviewQuestions()];
    if (list[index]) {
      list[index] = { ...list[index], marks: validMarks };
      this.reviewQuestions.set(list);
    }
  }

  deleteQuestion(index: number): void {
    const list = this.reviewQuestions().filter((_, i) => i !== index);
    this.reviewQuestions.set(list);
  }

  addQuestion(): void {
    const list = [...this.reviewQuestions()];
    list.push({
      question: 'New question text...',
      type: 'short_answer',
      marks: 2,
    });
    this.reviewQuestions.set(list);
  }

  backToPrompt(): void {
    this.stage.set('prompt');
  }

  onPublishAssignment(): void {
    if (!this.reviewTitle().trim()) {
      this.errorMessage.set('Assignment title cannot be empty.');
      return;
    }

    if (this.reviewQuestions().length === 0) {
      this.errorMessage.set('Please include at least one question in the assignment.');
      return;
    }

    this.stage.set('publishing');
    this.errorMessage.set(null);

    // Format questions into readable description for students
    let formattedDescription = `${this.reviewInstructions().trim()}\n\n---\n### Questions (Total Marks: ${this.calculatedTotalMarks()})\n`;
    this.reviewQuestions().forEach((q, idx) => {
      const typeLabel = this.questionTypes.find((t) => t.value === q.type)?.label || q.type;
      formattedDescription += `\n**Q${idx + 1} (${q.marks} Marks - ${typeLabel}):**\n${q.question}\n`;
    });

    const payload = {
      title: this.reviewTitle().trim(),
      className: this.reviewClassName().trim(),
      subject: this.reviewSubject().trim(),
      description: formattedDescription.trim(),
      attachments: '',
    };

    this.hwService.create(payload).subscribe({
      next: () => {
        this.stage.set('prompt');
        this.successMessage.set('AI Assignment published successfully to your classroom!');
        this.promptForm.patchValue({ prompt: '' });
        this.assignmentPublished.emit();

        setTimeout(() => {
          this.successMessage.set(null);
        }, 5000);
      },
      error: (err) => {
        this.stage.set('review');
        this.errorMessage.set(
          err?.error?.error || 'Failed to publish assignment. Please try again.'
        );
      },
    });
  }

  dismissAlert(): void {
    this.errorMessage.set(null);
    this.successMessage.set(null);
  }
}
