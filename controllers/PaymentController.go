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

type PaymentInput struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	AdvertiseID *uint   `json:"advertise_id"`
	VocationID  *uint   `json:"vocation_id"`
	UserID      uint    `json:"user_id"`
	DueDate     string  `json:"due_date"`
}

// GetAllPayment godoc
// @Summary Get all Payment.
// @Description Get a list of Payment.
// @Tags Payment
// @Produce json
// @Success 200 {object} []models.Payment
// @Router /payment [get]
func GetAllPayment(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var Payments []models.Payment
	db.Find(&Payments)

	c.JSON(http.StatusOK, Payments)
}

// CreatePayment godoc
// @Summary Create New Payment.
// @Description Creating a new Payment.
// @Tags Payment
// @Param Body body PaymentInput true "the body to create a new Payment"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Payment
// @Router /payment [post]
func CreatePayment(c *gin.Context) {

	//Container Time
	// now := time.Now()
	// yyyy, mm, dd := now.Date()

	// Validate input
	var input PaymentInput
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
	Payment := models.Payment{
		UserID:      userID,
		VocationID:  input.VocationID,
		AdvertiseID: input.AdvertiseID,
		Amount:      input.Amount,
		Currency:    input.Currency,
		DueDate:     input.DueDate,
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&Payment)

	c.JSON(http.StatusOK, Payment)
}

// GetPaymentById godoc
// @Summary Get Payment.
// @Description Get an Payment by id.
// @Tags Payment
// @Produce json
// @Param id path string true "Payment id"
// @Success 200 {object} models.Payment
// @Router /payment/{id} [get]
func GetPaymentById(c *gin.Context) { // Get model if exist
	var payment models.Payment

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?",
		c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

// UpdatePayment godoc
// @Summary Update Payment.
// @Description Update Payment by id.
// @Tags Payment
// @Produce json
// @Param id path string true "Payment id"
// @Param Body body PaymentInput true "the body to update Payment"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Payment
// @Router /payment/{id} [patch]
func UpdatePayment(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var payments models.Payment
	if err := db.Where("id = ?", c.Param("id")).First(&payments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input PaymentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Payment
	updatedInput.Amount = input.Amount
	updatedInput.Currency = input.Currency
	updatedInput.AdvertiseID = input.AdvertiseID
	updatedInput.VocationID = input.VocationID
	updatedInput.UserID = input.UserID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&payments).Updates(updatedInput)

	c.JSON(http.StatusOK, payments)
}

// DeletePayment godoc
// @Summary Delete one Payment.
// @Description Delete a Payment by id.
// @Tags Payment
// @Produce json
// @Param id path string true "Payment id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /payment/{id} [delete]
func DeletePayment(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var payments models.Payment
	if err := db.Where("id = ?", c.Param("id")).First(&payments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&payments)

	c.JSON(http.StatusOK, true)
}

// GetAllPaymentFromCurrentUser godoc
// @Summary Get Struct Payment from Curent User.
// @Description Get All Struct Payment from Curent User.
// @Tags Payment
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []models.Payment
// @Router /payment-current-user [get]
func GetAllPaymentFromCurrentUser(c *gin.Context) { // Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var userID, err = token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var payments []models.Payment
	if err := db.Where("user_id = ?", userID).Preload(clause.Associations).Find(&payments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}
