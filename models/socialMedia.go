package models

import (
	"time"
)

type SocialMedia struct{
	ID uint `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `json:"name" gorm:"type:varchar(50);notNull"`
	SocialMediaUrl string `json:"social_media_url" gorm:"type:text;notNull"`
	UserID uint `json:"user_id" gorm:"notNull"`
	User User
	CreatedAt time.Time `json:"created_at" gorm:"notNull"`
	UpdatedAt time.Time `json:"updated_at" gorm:"notNull"`
}

type RequestSocialMedia struct{
	Name string `json:"name" form:"name" valid:"required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_medial_url" valid:"required"`
}

type SocialMediaIncludeUser struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserID uint `json:"user_id"`
	User struct {
		ID string `json:"id"`
		Username string `json:"username"`
	} 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}