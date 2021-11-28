package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `json:"title" form:"title" valid:"required~Title of your photo is required"`
	Caption  string `json:"caption" form:"caption" valid:"required~Caption of your photo is required"`
	PhotoUrl string `json:"photo_url" form:"photo_url" valid:"required~PhotoUrl of your photo is required"`
	UserID   uint
	// User     *User
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
