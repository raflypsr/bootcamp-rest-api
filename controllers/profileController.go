package controllers

import (
	"net/http"
	"time"

	"fp-super-bootcamp-go/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type profileInput struct {
	Umur         int    `json:"umur" binding:"required"`
	NamaLengkap  string `json:"nama_lengkap" binding:"required"`
	JenisKelamin string `json:"jenis_kelamin" binding:"required"`
	UserID       uint   `json:"user_id" binding:"required"`
}

// GetAllUser godoc
// @Summary Get All profile.
// @Description Get a list of profile.
// @Tags profile
// @Produce json
// @Success 200 {object} []models.Profile
// @Router /profile [get]
func GetAllProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var profiles []models.Profile
	db.Find(&profiles)

	c.JSON(http.StatusOK, gin.H{"data": profiles})
}

// CreateProfile godoc
// @Summary Create New profile.
// @Description Creating a new Profile.
// @Tags profile
// @Param Body body profileInput true "the body to create a new Profile"
// @Produce json
// @Success 200 {object} models.Profile
// @Router /profile [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func CreateProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input profileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile := models.Profile{Umur: input.Umur, NamaLengkap: input.NamaLengkap, JenisKelamin: input.JenisKelamin, UserID: input.UserID}
	db.Create(&profile)

	c.JSON(http.StatusOK, gin.H{"data": profile})
}

// GetProfileByUserId godoc
// @Summary Get profile by user id.
// @Description Get a Profile by user id.
// @Tags profile
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} models.Profile
// @Router /user/{id}/profile [get]
func GetProfileByUserId(c *gin.Context) { // Get model if exist
	var profile models.Profile
	userID := c.Param("id")

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": profile})
}

// UpdateProfile godoc
// @Summary Update profile.
// @Description Update Profile by id.
// @Tags profile
// @Produce json
// @Param id path string true "Profile id"
// @Param Body body profileInput true "the body to update profile"
// @Success 200 {object} models.Profile
// @Router /profile/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var profile models.Profile
	if err := db.Where("id = ?", c.Param("id")).First(&profile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input profileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Profile
	updatedInput.Umur = input.Umur
	updatedInput.NamaLengkap = input.NamaLengkap
	updatedInput.JenisKelamin = input.JenisKelamin
	updatedInput.UpdatedAt = time.Now()

	db.Model(&profile).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": profile})
}

// DeleteProfile godoc
// @Summary Delete profile.
// @Description Delete a Profile by id.
// @Tags profile
// @Produce json
// @Param id path string true "Profile id"
// @Success 200 {object} map[string]boolean
// @Router /profile/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var profile models.Profile
	if err := db.Where("id = ?", c.Param("id")).First(&profile).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&profile)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
