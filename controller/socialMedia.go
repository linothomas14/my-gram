package controller

import (
	"my-gram/database"
	"my-gram/helpers"
	"my-gram/models"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func GetSocialMedias(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	socialMedias := make([]models.SocialMediaIncludeUser, 0)
	rows, err := db.Table("social_media").Select(`social_media.id, social_media.name, social_media.social_media_url, 
		social_media.user_id, users.id, users.username, social_media.created_at, social_media.updated_at`).
		Joins("JOIN users on users.id = social_media.user_id").Where("user_id = ?", userId).Rows()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	for rows.Next() {
		socialMedia := models.SocialMediaIncludeUser{}
		err := rows.Scan(&socialMedia.ID, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserID,
			&socialMedia.User.ID, &socialMedia.User.Username, &socialMedia.CreatedAt, &socialMedia.UpdatedAt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		socialMedias = append(socialMedias, socialMedia)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   socialMedias,
	})
}

func StoreSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	requestSocialMedia := models.RequestSocialMedia{}
	userId := uint(userData["id"].(float64))

	contentType := helpers.GetContentType(ctx)
	if contentType == appJSON {
		ctx.ShouldBindJSON(&requestSocialMedia)
	} else {
		ctx.ShouldBind(&requestSocialMedia)
	}

	_, err := govalidator.ValidateStruct(requestSocialMedia)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	socialMedia := models.SocialMedia{
		Name:           requestSocialMedia.Name,
		SocialMediaUrl: requestSocialMedia.SocialMediaUrl,
		UserID:         userId,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	err = db.Create(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": gin.H{
			"id":                socialMedia.ID,
			"name":              socialMedia.Name,
			"social_medial_url": socialMedia.SocialMediaUrl,
			"user_id":           socialMedia.UserID,
			"created_at":        socialMedia.CreatedAt,
		},
	})
}

func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	requestSocialMedia := models.RequestSocialMedia{}
	socialMediaId, _ := strconv.ParseUint(ctx.Param("socialMediaId"), 10, 64)
	userId := uint(userData["id"].(float64))

	contentType := helpers.GetContentType(ctx)
	if contentType == appJSON {
		ctx.ShouldBindJSON(&requestSocialMedia)
	} else {
		ctx.ShouldBind(&requestSocialMedia)
	}

	_, err := govalidator.ValidateStruct(requestSocialMedia)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	socialMedia := models.SocialMedia{}
	err = db.First(&socialMedia, "user_id = ? AND id = ?", userId, socialMediaId).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}
	
	socialMedia = models.SocialMedia{
		ID:             uint(socialMediaId),
		Name:           requestSocialMedia.Name,
		SocialMediaUrl: requestSocialMedia.SocialMediaUrl,
		UpdatedAt:      time.Now(),
	}
	err = db.Model(&socialMedia).Select("name", "social_media_url", "updated_at").Updates(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"id":                socialMedia.ID,
			"name":              socialMedia.Name,
			"social_medial_url": socialMedia.SocialMediaUrl,
			"user_id":           userId,
			"updated_at":        socialMedia.UpdatedAt,
		},
	})
}

func DeleteSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	socialMediaId, _ := strconv.ParseUint(ctx.Param("socialMediaId"), 10, 64)
	socialMedia := models.SocialMedia{}
	err := db.First(&socialMedia, "user_id = ? AND id = ?", userId, socialMediaId).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}

	err = db.Delete(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"message": "Your social media has been successfully deleted",
		},
	})
}
