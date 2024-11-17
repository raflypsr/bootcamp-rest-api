package controllers

import (
	"net/http"

	"fp-super-bootcamp-go/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type likeInput struct {
	UserID   uint `json:"user_id" binding:"required"`
	ReviewID uint `json:"review_id" binding:"required"`
}

// GetAllLike godoc
// @Summary Get all like.
// @Description Get a list of like.
// @Tags like
// @Produce json
// @Success 200 {object} []models.Likes
// @Router /like [get]
func GetAllLike(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var likes []models.Likes
	db.Find(&likes)

	c.JSON(http.StatusOK, gin.H{"data": likes})
}

// CreateLike godoc
// @Summary Create New like.
// @Description Creating a new Like.
// @Tags like
// @Param Body body likeInput true "the body to create a new Like"
// @Produce json
// @Success 201 {object} models.Likes
// @Router /like [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func CreateLike(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input likeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	like := models.Likes{UserID: input.UserID, ReviewID: input.ReviewID}
	db.Create(&like)

	c.JSON(http.StatusOK, gin.H{"data": like})
}

// DeleteLikegodoc
// @Summary Delete one like.
// @Description Delete a Like by id.
// @Tags like
// @Produce json
// @Param id path string true "Like id"
// @Success 200 {object} map[string]boolean
// @Router /like/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteLike(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var like models.Likes
	if err := db.Where("id = ?", c.Param("id")).First(&like).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&like)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
