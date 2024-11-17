package controllers

import (
	"net/http"
	"time"

	"fp-super-bootcamp-go/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type bookInput struct {
	Judul       string `json:"judul" binding:"required"`
	Penulis     string `json:"penulis" binding:"required"`
	Penerbit    string `json:"penerbit" binding:"required"`
	TahunTerbit int    `json:"tahun_terbit" binding:"required"`
}

// GetAllBook godoc
// @Summary Get all book.
// @Description Get a list of book.
// @Tags book
// @Produce json
// @Success 200 {object} []models.Books
// @Router /book [get]
func GetAllBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []models.Books
	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// CreateBook godoc
// @Summary Create New Book.
// @Description Creating a new Book.
// @Tags book
// @Param Body body bookInput true "the body to create a new Book"
// @Produce json
// @Success 201 {object} models.Books
// @Router /book [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func CreateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input bookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.Books{Judul: input.Judul, Penulis: input.Penulis, Penerbit: input.Penerbit, TahunTerbit: input.TahunTerbit}
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GetBookById godoc
// @Summary Get Book.
// @Description Get a book by id.
// @Tags book
// @Produce json
// @Param id path string true "Book id"
// @Success 200 {object} models.Books
// @Router /book/{id} [get]
func GetBookById(c *gin.Context) {
	var book models.Books

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook godoc
// @Summary Update Book.
// @Description Update Book by id.
// @Tags book
// @Produce json
// @Param id path string true "Book id"
// @Param Body body bookInput true "the body to update book"
// @Success 200 {object} models.Books
// @Router /book/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateBook(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var book models.Books
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input bookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Books
	updatedInput.Judul = input.Judul
	updatedInput.Penulis = input.Penulis
	updatedInput.Penerbit = input.Penerbit
	updatedInput.TahunTerbit = input.TahunTerbit
	updatedInput.UpdatedAt = time.Now()

	db.Model(&book).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook godoc
// @Summary Delete book.
// @Description Delete a Book by id.
// @Tags book
// @Produce json
// @Param id path string true "Book id"
// @Success 200 {object} map[string]boolean
// @Router /book/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Books
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
