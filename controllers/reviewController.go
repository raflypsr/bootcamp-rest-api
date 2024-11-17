package controllers

import (
	"net/http"
	"time"

	"fp-super-bootcamp-go/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type reviewInput struct {
	Deskripsi string `json:"deskripsi" binding:"required"`
	Judul string `json:"judul"`
	UserID    uint   `json:"user_id" binding:"required"`
	BookID    uint   `json:"book_id" binding:"required"`
}

// GetAllReview godoc
// @Summary Get all review.
// @Description Get a list of review.
// @Tags review
// @Produce json
// @Success 200 {object} []models.Reviews
// @Router /review [get]
func GetAllReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var reviews []models.Reviews
	db.Find(&reviews)

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// CreateBook godoc
// @Summary Create New review.
// @Description Creating a new Review.
// @Tags review
// @Param Body body reviewInput true "the body to create a new Review"
// @Produce json
// @Success 201 {object} models.Reviews
// @Router /review [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func CreateReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input reviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := models.Reviews{Deskripsi: input.Deskripsi, Judul: input.Judul, UserID: input.UserID, BookID: input.BookID}
	db.Create(&review)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// GetReviewById godoc
// @Summary Get review.
// @Description Get a Review by id.
// @Tags review
// @Produce json
// @Param id path string true "Review id"
// @Success 200 {object} models.Reviews
// @Router /review/{id} [get]
func GetReviewById(c *gin.Context) {
	var review models.Reviews

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// UpdateReview godoc
// @Summary Update review.
// @Description Update Review by id.
// @Tags review
// @Produce json
// @Param id path string true "Review id"
// @Param Body body reviewInput true "the body to update review"
// @Success 200 {object} models.Reviews
// @Router /review/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateReview(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var review models.Reviews
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input reviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Reviews
	updatedInput.Judul = input.Judul
	updatedInput.Deskripsi = input.Deskripsi
	updatedInput.UserID = input.UserID
	updatedInput.BookID = input.BookID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&review).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// DeleteReview godoc
// @Summary Delete review.
// @Description Delete a Review by id.
// @Tags review
// @Produce json
// @Param id path string true "Review id"
// @Success 200 {object} map[string]boolean
// @Router /review/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var review models.Reviews
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&review)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
