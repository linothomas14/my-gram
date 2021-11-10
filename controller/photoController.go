package controller

import (
	"my-gram/database"
	"my-gram/helpers"
	"my-gram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PostPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = UserID
	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": map[string]interface{}{
			"id":         Photo.ID,
			"title":      Photo.Title,
			"caption":    Photo.Caption,
			"photo_url":  Photo.PhotoUrl,
			"user_id":    Photo.UserID,
			"created_at": Photo.CreatedAt,
		},
	})
}

func ReadAllPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	var Photo []models.Photo
	UserID := uint(userData["id"].(float64))

	err := db.Debug().Where("user_id = ?", UserID).Find(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   Photo,
	})
}

// func UpdatePhoto(c *gin.Context){
// 	db := database.GetDB()
// 	contentType := helpers.GetContentType(c)

// }
