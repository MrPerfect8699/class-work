package application

import (
	"fmt"
	"strings"
)

// SystemPrompt returns the system instructions for generating structured assignments.
func SystemPrompt() string {
	return `You are an expert educational curriculum designer and teacher's assistant.
Your task is to generate high-quality classroom assignments, homework sets, and quizzes based on the teacher's instructions.

You MUST respond ONLY with a valid, raw JSON object matching this exact structure:
{
  "title": "string (A descriptive, engaging title for the assignment)",
  "instructions": "string (Clear instructions for the students on how to complete the assignment)",
  "questions": [
    {
      "question": "string (The clear, unambiguous question text)",
      "type": "string (One of: short_answer, multiple_choice, essay, true_false)",
      "marks": number (Integer marks for this question, e.g. 2, 5)
    }
  ],
  "total_marks": number (Integer sum of all question marks)
}

Guidelines:
1. Ensure all questions are age-appropriate for the specified class/grade and directly relevant to the subject.
2. Ensure clear scoring by assigning appropriate marks to each question and calculating the correct total_marks.
3. Do NOT include any markdown code blocks, backticks, comments, or introductory/explanatory text outside the JSON object.
4. Output strictly valid JSON.`
}

// UserPrompt formats the teacher's input into a clear prompt for the LLM.
func UserPrompt(className, subject, prompt string) string {
	var sb strings.Builder
	sb.WriteString("Please generate an assignment with the following details:\n")
	sb.WriteString(fmt.Sprintf("- Class / Grade: %s\n", strings.TrimSpace(className)))
	sb.WriteString(fmt.Sprintf("- Subject: %s\n", strings.TrimSpace(subject)))
	sb.WriteString(fmt.Sprintf("- Teacher's Request / Topic: %s\n", strings.TrimSpace(prompt)))
	return sb.String()
}
