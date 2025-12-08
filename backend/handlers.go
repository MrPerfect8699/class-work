package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if email already exists
	var existing Teacher
	err := MongoDB.Collection("teachers").FindOne(ctx, bson.M{"email": body.Email}).Decode(&existing)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already taken"})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	teacher := Teacher{
		ID:        primitive.NewObjectID(),
		Name:      body.Name,
		Email:     body.Email,
		Password:  string(hashed),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := MongoDB.Collection("teachers").InsertOne(ctx, teacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create teacher"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var teacher Teacher
	err := MongoDB.Collection("teachers").FindOne(ctx, bson.M{"email": body.Email}).Decode(&teacher)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(teacher.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-secret-key" // default for local development
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"teacher_id": teacher.ID.Hex(),
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
		TeacherID   string `json:"teacherId"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacherID, err := primitive.ObjectIDFromHex(body.TeacherID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid teacher ID"})
		return
	}

	homework := Homework{
		ID:          primitive.NewObjectID(),
		Title:       body.Title,
		Description: body.Description,
		ClassName:   body.ClassName,
		Subject:     body.Subject,
		TeacherID:   teacherID,
		Attachments: body.Attachments,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := MongoDB.Collection("homework").InsertOne(ctx, homework)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create homework"})
		return
	}

	homework.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, homework)
}

// List homeworks for a teacher
func ListHomeworks(c *gin.Context) {
	teacherID := c.Query("teacherId")

	objID, err := primitive.ObjectIDFromHex(teacherID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid teacher ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := MongoDB.Collection("homework").Find(ctx, bson.M{"teacherId": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch homework"})
		return
	}
	defer cursor.Close(ctx)

	var homeworks []Homework
	if err = cursor.All(ctx, &homeworks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not decode homework"})
		return
	}

	if homeworks == nil {
		homeworks = []Homework{}
	}

	c.JSON(http.StatusOK, homeworks)
}
