package main

import (
    "net/http"
    "os"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

// Register teacher
func Register(c *gin.Context) {
    var body struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
    teacher := Teacher{Name: body.Name, Email: body.Email, Password: string(hashed)}
    if err := DB.Create(&teacher).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "email maybe taken"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"id": teacher.ID})
}

// Login
func Login(c *gin.Context) {
    var body struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var t Teacher
    if DB.Where("email = ?", body.Email).First(&t).RecordNotFound() {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid creds"})
        return
    }
    if err := bcrypt.CompareHashAndPassword([]byte(t.Password), []byte(body.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid creds"})
        return
    }
    secret := os.Getenv("JWT_SECRET")
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "teacher_id": t.ID,
        "exp":        time.Now().Add(72 * time.Hour).Unix(),
    })
    s, _ := token.SignedString([]byte(secret))
    c.JSON(http.StatusOK, gin.H{"token": s})
}

// Create homework
func CreateHomework(c *gin.Context) {
    var body struct {
        Title       string `json:"title"`
        Description string `json:"description"`
        ClassName   string `json:"className"`
        Subject     string `json:"subject"`
        Attachments string `json:"attachments"`
        TeacherID   uint   `json:"teacherId"`
    }
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    h := Homework{
        Title:       body.Title,
        Description: body.Description,
        ClassName:   body.ClassName,
        Subject:     body.Subject,
        Attachments: body.Attachments,
        TeacherID:   body.TeacherID,
    }
    if err := DB.Create(&h).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create"})
        return
    }
    c.JSON(http.StatusCreated, h)
}

// List homeworks for a teacher
func ListHomeworks(c *gin.Context) {
    teacherID := c.Query("teacherId")
    var list []Homework
    DB.Where("teacher_id = ?", teacherID).Order("created_at desc").Find(&list)
    c.JSON(http.StatusOK, list)
}
