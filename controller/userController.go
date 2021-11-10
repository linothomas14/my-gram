package controller

import (
	"my-gram/database"
	"my-gram/helpers"
	"my-gram/models"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func IndexHandler(c *gin.Context) {
	c.String(200, "hello")
}

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"error":   "Bad Request",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": map[string]interface{}{
			"age":      User.Age,
			"email":    User.Email,
			"id":       User.ID,
			"username": User.Username,
		},
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password
	err := db.Debug().Where("email=?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"error":   "Unauthorized",
			"message": "Invalid email/password",
			"data":    nil,
		})
		return
	}
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {

		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"error":   "Unauthorized",
			"message": "Invalid email/password",
			"data":    nil,
		})
		return
	}
	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

func UserUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	UserID := uint(userData["id"].(float64))

	db.Where("id = ?", UserID).First(&User)

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	User.ID = UserID

	_, errUpdate := govalidator.ValidateStruct(User)

	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"error":   "Update Not Valid",
			"message": errUpdate,
			"data":    nil,
		})
		return
	}

	err := db.Debug().Model(&User).Updates(models.User{Email: User.Email, Username: User.Username}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"error":   "Bad Request",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": map[string]interface{}{
			"age":        User.Age,
			"email":      User.Email,
			"id":         User.ID,
			"username":   User.Username,
			"updated_at": User.UpdatedAt,
		},
	})
}

func UserDelete(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	UserID := uint(userData["id"].(float64))

	err := db.Where("id = ?", UserID).Delete(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err,
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": map[string]interface{}{
			"message": "Your account has been successfully deleted",
		},
	})
}
