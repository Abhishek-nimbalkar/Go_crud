package controller

import (
	"log"
	"net/http"

	"example.com/crud/configs/db"
	"example.com/crud/pkg/models"
	"github.com/gin-gonic/gin"
)

type User struct {
	FirstName string `json:"FirstName" binding:"required"`
	LastName  string `json:"LastName" binding:"required"`
	Email     string `json:"Email" binding:"required"`
	Password  string `json:"Password" binding:"required"`
}

func ServerHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}

func CreateUser(c *gin.Context) {
	user := User{} // Assuming this is the User struct from your models package
	if err := c.BindJSON(&user); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received user data: %+v", user)

	// Check if a user with the same email already exists
	var existingUser models.User
	if err := db.DB.Where("email =?", user.Email).First(&existingUser).Error; err == nil && existingUser.ID != 0 {
		// Handle the case where a user with the same email already exists
		// For example, you might want to return an error response
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "A user with this email already exists."})
		return
	}

	newUser := models.User{ // This is the User struct from your models package
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}
	x := db.DB.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{"data": newUser, "error": x.Error})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUser(c *gin.Context) {
	userId:= c.Param("id")
	var user models.User
	if err := db.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := db.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	var input User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser := models.User{}

	db.DB.Model(&user).Updates(&updatedUser)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
