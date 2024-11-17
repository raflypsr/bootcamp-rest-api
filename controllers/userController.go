package controllers

import (
	"net/http"
	"time"

	"fp-super-bootcamp-go/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// GetAllUser godoc
// @Summary Get All user.
// @Description Get a list of user.
// @Tags user
// @Produce json
// @Success 200 {object} []models.Users
// @Router /user [get]
func GetAllUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.Users
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUserById godoc
// @Summary Get user.
// @Description Get an User by id.
// @Tags user
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} models.Users
// @Router /user/{id} [get]
func GetUserById(c *gin.Context) { // Get model if exist
	var user models.Users

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateUser godoc
// @Summary Update user.
// @Description Update User by id.
// @Tags user
// @Produce json
// @Param id path string true "User id"
// @Param Body body userInput true "the body to update user"
// @Success 200 {object} models.Users
// @Router /user/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var user models.Users
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input userInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create updated input with hashed password if password is provided
	var updatedInput models.Users
	updatedInput.Username = input.Username
	updatedInput.Email = input.Email
	updatedInput.Role = input.Role

	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		updatedInput.Password = string(hashedPassword)
	} else {
		updatedInput.Password = user.Password // Retain the old password if new password is not provided
	}

	updatedInput.UpdatedAt = time.Now()

	db.Model(&user).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser godoc
// @Summary Delete  user.
// @Description Delete a User by id.
// @Tags user
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} map[string]boolean
// @Router /user/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteUser(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var user models.Users
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
