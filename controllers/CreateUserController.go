package controllers

// import (
// 	"net/http"
// 	"time"

// 	"api-gin/models"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

type CreateUser struct {
	Title               string `json:"title"`
	Year                int    `json:"year"`
	AgeRatingCategoryID uint   `json:"age_rating_category_id"`
}
