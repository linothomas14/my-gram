package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// type Photo struct {
// 	GormModel
// 	Title    string `gorm:"primaryKey;autoIncrement:true" json:"title" form:"title" valid:"required~Title of your photo is required"`
// 	Caption  string `gorm:"notNull" json:"caption" form:"caption" valid:"required~Caption of your photo is required"`
// 	PhotoUrl string `gorm:"notNull" json:"photo_url" form:"photo_url" valid:"required~PhotoUrl of your photo is required"`
// 	UserID   uint   `gorm:"notNull"`
// 	// User     *User
// }
type Photo struct {
	GormModel
	Title    string `json:"title" form:"title" valid:"required~Title of your photo is required"`
	Caption  string `json:"caption" form:"caption" valid:"required~Caption of your photo is required"`
	PhotoUrl string `json:"photo_url" form:"photo_url" valid:"required~PhotoUrl of your photo is required"`
	UserID   uint
	// User     *User
}

type PhotoIncludeUserData struct {
	Id         uint       `json:"id"`
	Title      string     `json:"title" `
	Caption    string     `json:"caption" `
	PhotoUrl   string     `json:"photo_url" `
	UserID     uint       `json:"user_id"`
	Created_at *time.Time `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
	User       struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}
}

func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errVal := govalidator.ValidateStruct(u)
	if errVal != nil {
		err = errVal
		return
	}
	err = nil

	return
}
