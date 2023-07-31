package controllers

import (
	"2hf/models"
	"net/http"
	"time"

	"2hf/utils/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AdvertiseInput struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	UserID     uint   `json:"user_id"`
	VocationID uint   `json:"vocation_id"`
	Name       string `json:"name"`
	Duration   string `json:"duration"`
	Price      string `json:"price"`
	Content    string `json:"content"`
}

// GetAllAdvertise godoc
// @Summary Get all Advertise.
// @Description Get a list of Advertise.
// @Tags Advertise
// @Produce json
// @Success 200 {object} []models.Advertise
// @Router /advertise [get]
func GetAllAdvertise(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var advertises []models.Advertise
	db.Find(&advertises)

	c.JSON(http.StatusOK, advertises)
}

// CreateAdvertise godoc
// @Summary Create New Advertise.
// @Description Creating a new Advertise.
// @Tags Advertise
// @Param Body body AdvertiseInput true "the body to create a new Advertise"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Advertise
// @Router /advertise [post]
func CreateAdvertise(c *gin.Context) {
	// Validate input
	var input AdvertiseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// UserID input
	var userID, err = token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Cart
	advertises := models.Advertise{
		UserID:     userID,
		VocationID: input.VocationID,
		Name:       input.Name,
		Duration:   input.Duration,
		Price:      input.Price,
		Content:    input.Content,
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&advertises)

	c.JSON(http.StatusOK, advertises)
}

// GetAdvertiseById godoc
// @Summary Get Advertise.
// @Description Get an Advertise by id.
// @Tags Advertise
// @Produce json
// @Param id path string true "Advertise id"
// @Success 200 {object} models.Advertise
// @Router /advertise/{id} [get]
func GetAdvertiseById(c *gin.Context) { // Get model if exist
	var advertises models.Advertise

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?",
		c.Param("id")).First(&advertises).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, advertises)
}

// UpdateAdvertise godoc
// @Summary Update Advertise.
// @Description Update Advertise by id.
// @Tags Advertise
// @Produce json
// @Param id path string true "Advertise id"
// @Param Body body AdvertiseInput true "the body to update Advertise"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Advertise
// @Router /advertise/{id} [patch]
func UpdateAdvertise(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var advertises models.Advertise
	if err := db.Where("id = ?", c.Param("id")).First(&advertises).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input AdvertiseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Advertise
	updatedInput.Name = input.Name
	updatedInput.Duration = input.Duration
	updatedInput.Price = input.Price
	updatedInput.Content = input.Content
	updatedInput.UpdatedAt = time.Now()

	db.Model(&advertises).Updates(updatedInput)

	c.JSON(http.StatusOK, advertises)
}

// DeleteAdvertise godoc
// @Summary Delete one Advertise.
// @Description Delete a Advertise by id.
// @Tags Advertise
// @Produce json
// @Param id path string true "Advertise id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /advertise/{id} [delete]
func DeleteAdvertise(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var advertises models.Advertise
	if err := db.Where("id = ?", c.Param("id")).First(&advertises).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&advertises)

	c.JSON(http.StatusOK, true)
}

// GetAllAdvertiseFromCurrentUser godoc
// @Summary Get Struct Advertise from Curent User.
// @Description Get All Struct Advertise from Curent User.
// @Tags Advertise
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Advertise
// @Router /advertise-current-user [get]
func GetAllAdvertiseFromCurrentUser(c *gin.Context) { // Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var userID, err = token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var advertises []models.Advertise
	if err := db.Where("user_id = ?", userID).Preload(clause.Associations).Find(&advertises).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, advertises)
}
