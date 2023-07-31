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

type VocationInput struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Comment     string `json:"comment"`
	Rating      string `json:"rating"`
	Phone       string `json:"phone"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Image       string `json:"image"`
	Price       string `json:"price"`
	UserID      uint   `json:"user_id"`
}

// GetAllVocation godoc
// @Summary Get all Vocation.
// @Description Get a list of Vocation.
// @Tags Vocation
// @Produce json
// @Success 200 {object} []models.Vocation
// @Router /vocation [get]
func GetAllVocation(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var Vocations []models.Vocation
	db.Find(&Vocations)

	c.JSON(http.StatusOK, Vocations)
}

// CreateVocation godoc
// @Summary Create New Vocation.
// @Description Creating a new Vocation.
// @Tags Vocation
// @Param Body body VocationInput true "the body to create a new Vocation"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Vocation
// @Router /vocation [post]
func CreateVocation(c *gin.Context) {
	// Validate input
	var input VocationInput
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
	Vocation := models.Vocation{
		Name:        input.Name,
		Description: input.Description,
		Comment:     input.Comment,
		Rating:      input.Rating,
		Phone:       input.Phone,
		Country:     input.Country,
		City:        input.City,
		Address:     input.Address,
		Image:       input.Image,
		Price:       input.Price,
		UserID:      userID,
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&Vocation)

	c.JSON(http.StatusOK, Vocation)
}

// GetVocationById godoc
// @Summary Get Vocation.
// @Description Get an Vocation by id.
// @Tags Vocation
// @Produce json
// @Param id path string true "Vocation id"
// @Success 200 {object} models.Vocation
// @Router /vocation/{id} [get]
func GetVocationById(c *gin.Context) { // Get model if exist
	var vocation models.Vocation

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?",
		c.Param("id")).First(&vocation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, vocation)
}

// UpdateVocation godoc
// @Summary Update Vocation.
// @Description Update Vocation by id.
// @Tags Vocation
// @Produce json
// @Param id path string true "Vocation id"
// @Param Body body VocationInput true "the body to update Vocation"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Vocation
// @Router /vocation/{id} [patch]
func UpdateVocation(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var vocation models.Vocation
	if err := db.Where("id = ?", c.Param("id")).First(&vocation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input VocationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Vocation
	updatedInput.Name = input.Name
	updatedInput.Description = input.Description
	updatedInput.Comment = input.Comment
	updatedInput.Rating = input.Rating
	updatedInput.Phone = input.Phone
	updatedInput.Country = input.Country
	updatedInput.City = input.City
	updatedInput.Address = input.Address
	updatedInput.Image = input.Image
	updatedInput.Price = input.Price

	updatedInput.UpdatedAt = time.Now()

	db.Model(&vocation).Updates(updatedInput)

	c.JSON(http.StatusOK, vocation)
}

// DeleteVocation godoc
// @Summary Delete one Vocation.
// @Description Delete a Vocation by id.
// @Tags Vocation
// @Produce json
// @Param id path string true "Vocation id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /vocation/{id} [delete]
func DeleteVocation(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var Vocation models.Vocation
	if err := db.Where("id = ?", c.Param("id")).First(&Vocation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&Vocation)

	c.JSON(http.StatusOK, true)
}

// GetAllVocationFromCurrentUser godoc
// @Summary Get Struct Vocation from Curent User.
// @Description Get All Struct Vocation from Curent User.
// @Tags Vocation
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Vocation
// @Router /vocation-current-user [get]
func GetAllVocationFromCurrentUser(c *gin.Context) { // Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var userID, err = token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var Vocations []models.Vocation
	if err := db.Where("user_id = ?", userID).Preload(clause.Associations).Find(&Vocations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, Vocations)
}
