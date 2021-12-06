package models

import (
	"my-gram/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string  `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required"`
	Email    string  `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format"`
	Password string  `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int     `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required" validate:"min=8,max=200,numeric"`
	Photos   []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	// cant set validation minimal age , we will fixed it later .. code --->> ,min=8~sorry are underage,numeric~Must enter a number
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	validate := validator.New()
	_, errVal := govalidator.ValidateStruct(u)
	errCreate := validate.Struct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	if errVal != nil {
		err = errVal
		return
	}
	err = nil
	u.Password = helpers.HassPass(u.Password)
	return
}
