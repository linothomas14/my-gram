package controller

import (
	"my-gram/database"
	"my-gram/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetSocialMedias (ctx *gin.Context){
	db := database.GetDB()
	socialMedias := make([]models.SocialMediaIncludeUser, 0)
	rows, err := db.Table("social_media").Select(`social_media.id, social_media.name, social_media.social_media_url, 
		social_media.user_id, users.id, users.username, social_media.created_at, social_media.updated_at`).
		Joins("JOIN users on users.id = social_media.user_id").Rows()
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
		"data": socialMedias,
	})
}

func StoreSocialMedia(ctx *gin.Context){
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	requestSocialMedia := models.RequestSocialMedia{}
	userId := uint(userData["id"].(float64))
	ctx.ShouldBindJSON(&requestSocialMedia)
	socialMedia := models.SocialMedia { 
		Name: requestSocialMedia.Name, 
		SocialMediaUrl: requestSocialMedia.SocialMediaUrl, 
		UserID : userId, 
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}
	err := db.Create(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error(),})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"id": socialMedia.ID,
		"name": socialMedia.Name,
		"social_medial_url": socialMedia.SocialMediaUrl,
		"user_id": socialMedia.UserID,
		"created_at": socialMedia.CreatedAt,
	})
}

func UpdateSocialMedia(ctx *gin.Context){
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	requestSocialMedia := models.RequestSocialMedia{}
	socialMediaId,_ := strconv.ParseUint(ctx.Param("socialMediaId"),10,64)
	userId := uint(userData["id"].(float64))
	ctx.ShouldBindJSON(&requestSocialMedia)
	socialMedia := models.SocialMedia { 
		ID: uint(socialMediaId),
		Name: requestSocialMedia.Name, 
		SocialMediaUrl: requestSocialMedia.SocialMediaUrl, 
		UpdatedAt : time.Now(),
	}
	err := db.Model(&socialMedia).Select("name", "social_media_url", "updated_at").Updates(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error(),})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": socialMedia.ID,
		"name": socialMedia.Name,
		"social_medial_url": socialMedia.SocialMediaUrl,
		"user_id": userId,
		"updated_at": socialMedia.UpdatedAt,
	})
}

func DeleteSocialMedia(ctx *gin.Context){
	db := database.GetDB()
	socialMediaId, _ := strconv.ParseUint(ctx.Param("socialMediaId"),10,64)
	socialMedia := &models.SocialMedia{ID:uint(socialMediaId)}
	err := db.Delete(&socialMedia).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error(),})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message" : "Your social media has been successfully deleted"})
}