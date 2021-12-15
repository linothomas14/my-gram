package controller

import (
	"my-gram/database"
	"my-gram/models"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetComments(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	comments := make([]models.CommentIncludeUserPhoto, 0)
	rows, err := db.Table("comments").Select(`comments.id, comments.user_id, comments.photo_id, comments.message, 
		comments.created_at, comments.updated_at, users.id, users.email, users.username, photos.id, photos.title, 
		photos.caption, photos.photo_url, photos.user_id`).Joins(`JOIN users ON users.id = comments.user_id 
		JOIN photos ON photos.id = comments.photo_id`).Where("comments.user_id = ?", userId).Rows()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	for rows.Next() {
		comment := models.CommentIncludeUserPhoto{}
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.PhotoID, &comment.Message, &comment.CreatedAt,
			&comment.UpdatedAt, &comment.User.ID, &comment.User.Email, &comment.User.Username, &comment.Photo.ID,
			&comment.Photo.Title, &comment.Photo.Caption, &comment.Photo.PhotoUrl, &comment.Photo.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		comments = append(comments, comment)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   comments})

}

func StoreComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	requestComment := models.RequestComment{}
	userId := uint(userData["id"].(float64))
	ctx.ShouldBindJSON(&requestComment)

	_, err := govalidator.ValidateStruct(requestComment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	comment := models.Comment{
		UserID:    userId,
		PhotoID:   requestComment.PhotoID,
		Message:   requestComment.Message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = db.Create(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": gin.H{

			"id":         comment.ID,
			"message":    comment.Message,
			"user_id":    comment.UserID,
			"photo_id":   comment.PhotoID,
			"created_at": comment.CreatedAt,
		}})
}

func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	requestComment := models.RequestComment{}
	commentId, _ := strconv.ParseUint(ctx.Param("commentId"), 10, 64)
	ctx.ShouldBindJSON(&requestComment)

	_, err := govalidator.ValidateStruct(requestComment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	comment := models.Comment{}
	err = db.First(&comment, "user_id = ? AND id = ?", userId, commentId).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}

	comment = models.Comment{
		ID:        uint(commentId),
		Message:   requestComment.Message,
		UpdatedAt: time.Now(),
	}

	err = db.Model(&comment).Select("message", "updated_at").Updates(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	photo := models.Photo{}
	err = db.First(&photo, "id = ?", userId).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": gin.H{
			"id": comment.ID,
			"title": photo.Title,
			"caption": photo.Caption,
			"photo_url": photo.PhotoUrl,
			"user_id" : userId,
			"updated_at": comment.UpdatedAt,
		},
	})
}

func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	commentId, _ := strconv.ParseUint(ctx.Param("commentId"), 10, 64)
	comment := models.Comment{}
	err := db.First(&comment, "user_id = ? AND id = ?", userId, commentId).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}
	comment = models.Comment{ID: uint(commentId)}
	err = db.Delete(&comment).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"message": "Your comment has been successfully deleted"}})
}
