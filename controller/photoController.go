package controller

import (
	"fmt"
	"log"
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
	UserID := uint(userData["id"].(float64))

	var Photo []models.Photo
	var User models.User
	User.ID = UserID
	Email := fmt.Sprintf("%v", userData["email"])

	err := db.Debug().Where("user_id = ?", UserID).Find(&Photo).Error
	db.Select("username").Find(&User)

	var Photos []models.PhotoIncludeUserData
	var TempPhoto models.PhotoIncludeUserData

	for i := range Photo {
		TempPhoto.Id = Photo[i].ID
		TempPhoto.Title = Photo[i].Title
		TempPhoto.Caption = Photo[i].Caption
		TempPhoto.PhotoUrl = Photo[i].PhotoUrl
		TempPhoto.UserID = Photo[i].UserID
		TempPhoto.Created_at = Photo[i].CreatedAt
		TempPhoto.Updated_at = Photo[i].UpdatedAt
		TempPhoto.User.Email = Email
		TempPhoto.User.Username = User.Username
		Photos = append(Photos, TempPhoto)

	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Photos,
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

	Photo.ID = uint(idTemp)

	err = db.First(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Cant find idPhoto(2)",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	result := db.Model(&Photo).Where("user_id = ?", UserID).Updates(&Photo)

	log.Println(Photo)
	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "No row affected",
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
		"message": "Your photo has been successfully deleted",
	})
}
