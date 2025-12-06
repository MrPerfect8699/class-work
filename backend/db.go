package main

import (
    "fmt"
    "log"
    "os"

    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL not set")
    }
    db, err := gorm.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("failed to connect db: %v", err)
    }
    db.AutoMigrate(&Teacher{}, &Homework{}, &Submission{})
    DB = db
    fmt.Println("Connected to DB")
}
