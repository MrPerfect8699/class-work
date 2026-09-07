import { Injectable } from '@angular/core';
import { GeneratedQuestion, Homework } from '../entities/models';

export interface WorksheetData {
  title: string;
  className: string;
  subject: string;
  instructions: string;
  questions: GeneratedQuestion[];
  totalMarks?: number;
}

@Injectable({ providedIn: 'root' })
export class ExportService {

  /**
   * Generates a printable student worksheet paper in a new window and triggers the browser print / Save as PDF dialog.
   */
  printWorksheet(data: WorksheetData): void {
    const totalMarks = data.totalMarks ?? data.questions.reduce((sum, q) => sum + (Number(q.marks) || 0), 0);

    const printWindow = window.open('', '_blank', 'width=850,height=900');
    if (!printWindow) {
      alert('Please allow popups to print/export the assignment worksheet.');
      return;
    }

    const questionItemsHtml = data.questions
      .map((q, idx) => {
        const typeLabel = this.formatTypeLabel(q.type);
        // Replace newlines in questions with <br> for MCQ choices or multi-line questions
        const formattedQuestion = this.escapeHtml(q.question).replace(/\n/g, '<br>');

        return `
          <div class="question-block">
            <div class="question-header">
              <span class="q-number">Q${idx + 1}.</span>
              <span class="q-meta">[${q.marks} Mark${q.marks > 1 ? 's' : ''} | ${typeLabel}]</span>
            </div>
            <div class="question-body">
              ${formattedQuestion}
            </div>
            ${q.type === 'essay' || q.type === 'short_answer' ? '<div class="answer-space"></div>' : ''}
          </div>
        `;
      })
      .join('');

    const htmlContent = `
      <!DOCTYPE html>
      <html lang="en">
      <head>
        <meta charset="UTF-8">
        <title>${this.escapeHtml(data.title)} - Worksheet</title>
        <style>
          @page {
            size: A4;
            margin: 18mm 15mm;
          }
          * {
            box-sizing: border-box;
            -webkit-print-color-adjust: exact;
            print-color-adjust: exact;
          }
          body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
            color: #1e293b;
            margin: 0;
            padding: 20px;
            font-size: 13.5pt;
            line-height: 1.5;
          }
          .worksheet-header {
            border-bottom: 2px solid #0f172a;
            padding-bottom: 12px;
            margin-bottom: 16px;
          }
          .school-title {
            text-align: center;
            font-size: 18pt;
            font-weight: 800;
            text-transform: uppercase;
            letter-spacing: 0.5px;
            color: #0f172a;
            margin: 0 0 6px;
          }
          .assignment-heading {
            text-align: center;
            font-size: 15pt;
            font-weight: 700;
            color: #334155;
            margin: 0 0 12px;
          }
          .meta-table {
            width: 100%;
            border-collapse: collapse;
            font-size: 11pt;
            margin-bottom: 8px;
          }
          .meta-table td {
            padding: 4px 6px;
          }
          .meta-label {
            font-weight: 600;
            color: #475569;
          }
          .fill-line {
            border-bottom: 1px dotted #64748b;
            display: inline-block;
            width: 160px;
          }
          .instructions-box {
            background-color: #f8fafc;
            border-left: 4px solid #3b82f6;
            padding: 10px 14px;
            font-size: 11pt;
            margin-bottom: 20px;
            color: #334155;
          }
          .instructions-label {
            font-weight: 700;
            color: #1e293b;
            margin-bottom: 4px;
          }
          .question-block {
            margin-bottom: 20px;
            page-break-inside: avoid;
          }
          .question-header {
            display: flex;
            justify-content: space-between;
            align-items: baseline;
            font-weight: 700;
            margin-bottom: 4px;
          }
          .q-number {
            font-size: 13pt;
            color: #0f172a;
          }
          .q-meta {
            font-size: 10pt;
            color: #64748b;
            font-weight: 600;
          }
          .question-body {
            font-size: 12pt;
            color: #1e293b;
            margin-left: 18px;
          }
          .answer-space {
            margin-top: 12px;
            margin-left: 18px;
            min-height: 40px;
            border-bottom: 1px dashed #cbd5e1;
          }
          @media print {
            body { padding: 0; }
            .no-print { display: none !important; }
          }
        </style>
      </head>
      <body>
        <div class="worksheet-header">
          <div class="school-title">ClassWork Academy</div>
          <div class="assignment-heading">${this.escapeHtml(data.title)}</div>
          
          <table class="meta-table">
            <tr>
              <td><span class="meta-label">Class:</span> <strong>${this.escapeHtml(data.className)}</strong></td>
              <td><span class="meta-label">Subject:</span> <strong>${this.escapeHtml(data.subject)}</strong></td>
              <td><span class="meta-label">Total Marks:</span> <strong>${totalMarks}</strong></td>
            </tr>
            <tr>
              <td colspan="2"><span class="meta-label">Student Name:</span> <span class="fill-line" style="width: 250px;"></span></td>
              <td><span class="meta-label">Date:</span> <span class="fill-line" style="width: 100px;"></span></td>
            </tr>
          </table>
        </div>

        ${
          data.instructions
            ? `
          <div class="instructions-box">
            <div class="instructions-label">General Instructions:</div>
            <div>${this.escapeHtml(data.instructions)}</div>
          </div>
        `
            : ''
        }

        <div class="questions-container">
          ${questionItemsHtml}
        </div>

        <script>
          window.onload = function() {
            setTimeout(function() {
              window.print();
            }, 300);
          };
        </script>
      </body>
      </html>
    `;

    printWindow.document.open();
    printWindow.document.write(htmlContent);
    printWindow.document.close();
  }

  /**
   * Downloads the assignment formatted as a plain text (.txt) document.
   */
  downloadText(data: WorksheetData): void {
    const totalMarks = data.totalMarks ?? data.questions.reduce((sum, q) => sum + (Number(q.marks) || 0), 0);

    let content = `================================================================================\n`;
    content += `CLASSWORK - ASSIGNMENT WORKSHEET\n`;
    content += `Title:       ${data.title}\n`;
    content += `Class:       ${data.className}\n`;
    content += `Subject:     ${data.subject}\n`;
    content += `Total Marks: ${totalMarks}\n`;
    content += `================================================================================\n\n`;

    if (data.instructions) {
      content += `INSTRUCTIONS:\n${data.instructions}\n\n`;
      content += `--------------------------------------------------------------------------------\n\n`;
    }

    data.questions.forEach((q, idx) => {
      const typeLabel = this.formatTypeLabel(q.type);
      content += `Q${idx + 1}. [${q.marks} Mark${q.marks > 1 ? 's' : ''} | ${typeLabel}]\n`;
      content += `${q.question}\n\n`;
    });

    const safeTitle = this.sanitizeFilename(data.title);
    this.triggerDownload(content, `${safeTitle}-worksheet.txt`, 'text/plain;charset=utf-8');
  }

  /**
   * Downloads the raw JSON data file.
   */
  downloadJson(data: any, filename: string): void {
    const jsonStr = JSON.stringify(data, null, 2);
    this.triggerDownload(jsonStr, `${this.sanitizeFilename(filename)}.json`, 'application/json;charset=utf-8');
  }

  /**
   * Exports an existing published homework record.
   */
  printHomework(hw: Homework): void {
    const parsedQuestions = this.extractQuestionsFromDescription(hw.description || '');

    this.printWorksheet({
      title: hw.title,
      className: hw.className || 'General',
      subject: hw.subject || 'General',
      instructions: parsedQuestions.instructions,
      questions: parsedQuestions.questions,
      totalMarks: parsedQuestions.totalMarks,
    });
  }

  downloadHomeworkText(hw: Homework): void {
    const parsedQuestions = this.extractQuestionsFromDescription(hw.description || '');

    this.downloadText({
      title: hw.title,
      className: hw.className || 'General',
      subject: hw.subject || 'General',
      instructions: parsedQuestions.instructions,
      questions: parsedQuestions.questions,
      totalMarks: parsedQuestions.totalMarks,
    });
  }

  // --- Helper Methods ---

  private formatTypeLabel(type: string): string {
    switch (type) {
      case 'multiple_choice':
        return 'Multiple Choice';
      case 'short_answer':
        return 'Short Answer';
      case 'essay':
        return 'Essay / Descriptive';
      case 'true_false':
        return 'True / False';
      default:
        return type || 'Question';
    }
  }

  private escapeHtml(str: string): string {
    return (str || '')
      .replace(/&/g, '&amp;')
      .replace(/</g, '&lt;')
      .replace(/>/g, '&gt;')
      .replace(/"/g, '&quot;')
      .replace(/'/g, '&#039;');
  }

  private sanitizeFilename(name: string): string {
    return (name || 'assignment')
      .toLowerCase()
      .replace(/[^a-z0-9_-]/gi, '_')
      .replace(/_+/g, '_')
      .substring(0, 50);
  }

  private triggerDownload(content: string, filename: string, mimeType: string): void {
    const blob = new Blob([content], { type: mimeType });
    const url = window.URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.download = filename;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    window.URL.revokeObjectURL(url);
  }

  private extractQuestionsFromDescription(description: string): {
    instructions: string;
    questions: GeneratedQuestion[];
    totalMarks: number;
  } {
    // If the description has our structured questions format, parse questions
    if (!description.includes('### Questions')) {
      return {
        instructions: description,
        questions: [{ question: description, type: 'General', marks: 10 }],
        totalMarks: 10,
      };
    }

    const parts = description.split('### Questions');
    const instructions = parts[0].replace('---', '').trim();
    const questionsBlock = parts[1] || '';

    const questions: GeneratedQuestion[] = [];
    const qMatches = questionsBlock.split(/\*\*Q\d+/g);

    for (let i = 1; i < qMatches.length; i++) {
      const chunk = qMatches[i];
      // Match marks and text
      const marksMatch = chunk.match(/\((\d+)\s*Marks/i);
      const marks = marksMatch ? parseInt(marksMatch[1], 10) : 2;

      // Extract text after header line
      const colonIdx = chunk.indexOf('):\n');
      const text = colonIdx !== -1 ? chunk.substring(colonIdx + 3).trim() : chunk.trim();

      questions.push({
        question: text,
        type: 'question',
        marks: marks,
      });
    }

    const totalMarks = questions.reduce((sum, q) => sum + q.marks, 0);

    return {
      instructions,
      questions: questions.length > 0 ? questions : [{ question: description, type: 'General', marks: 10 }],
      totalMarks: totalMarks || 10,
    };
  }
}
