package main

import "github.com/jinzhu/gorm"

type Teacher struct {
    gorm.Model
    Name     string `gorm:"size:100"`
    Email    string `gorm:"unique;size:100"`
    Password string `gorm:"size:255"`
}

type Homework struct {
    gorm.Model
    Title       string
    Description string `gorm:"type:text"`
    ClassName   string
    Subject     string
    TeacherID   uint
    Attachments string // store file path or URL, simple for MVP
}

type Submission struct {
    gorm.Model
    HomeworkID  uint
    StudentName string
    Completed   bool
    Notes       string
}
