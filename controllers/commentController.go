package controllers

import (
	"net/http"
	"time"

	"fp-super-bootcamp-go/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentInput struct {
	Deskripsi string `json:"deskripsi" binding:"required"`
	UserID    uint   `json:"user_id" binding:"required"`
	ReviewID  uint   `json:"review_id" binding:"required"`
}

// GetAllComment godoc
// @Summary Get all comment.
// @Description Get a list of comment.
// @Tags comment
// @Produce json
// @Success 200 {object} []models.Comments
// @Router /comment [get]
func GetAllComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var comments []models.Comments
	db.Find(&comments)

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// CreateComment godoc
// @Summary Create New Comment.
// @Description Creating a new Comment.
// @Tags comment
// @Param Body body commentInput true "the body to create a new Comment"
// @Produce json
// @Success 201 {object} models.Comments
// @Router /comment [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func CreateComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input commentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := models.Comments{Deskripsi: input.Deskripsi, UserID: input.UserID, ReviewID: input.ReviewID}
	db.Create(&review)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// UpdateComment godoc
// @Summary Update Comment.
// @Description Update Comment by id.
// @Tags comment
// @Produce json
// @Param id path string true "Comment id"
// @Param Body body commentInput true "the body to update comment"
// @Success 200 {object} models.Comments
// @Router /comment/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateComment(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var comment models.Comments
	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input commentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Comments
	updatedInput.Deskripsi = input.Deskripsi
	updatedInput.UserID = input.UserID
	updatedInput.ReviewID = input.ReviewID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&comment).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// DeleteCommentgodoc
// @Summary Delete comment.
// @Description Delete a Comment by id.
// @Tags comment
// @Produce json
// @Param id path string true "Comment id"
// @Success 200 {object} map[string]boolean
// @Router /comment/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var comment models.Comments
	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&comment)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
