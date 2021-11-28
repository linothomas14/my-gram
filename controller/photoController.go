package controller

import (
	"my-gram/database"
	"my-gram/helpers"
	"my-gram/models"
	"net/http"
	"strconv"

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

			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{

		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
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
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Photo,
	})
}
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	UserID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Photo := models.Photo{}

	idTemp, err := strconv.ParseUint(c.Param("photoId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Cant find idPhoto(1)",
		})
		return
	}
	var id uint = uint(idTemp)

	Photo.ID = id

	err = db.First(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Cant find idPhoto",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	// Photo.UserID = UserID

	result := db.Model(&Photo).Where("user_id = ?", UserID).Updates(Photo)
	// log.Println(Photo)
	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Cant find idPhoto",
		})
		return
	}
	c.JSON(http.StatusOK, Photo)
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	UserID := uint(userData["id"].(float64))

	Photo := models.Photo{}
	id := c.Param("photoId")

	result := db.Debug().Where("user_id = ?", UserID).Delete(&Photo, id)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Cant find idPhoto",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"user_id": userData,
		"message": "Your photo has been successfully deleted",
	})
}
