package models

import "time"

type Comment struct{
	ID uint `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserID uint `json:"user_id" gorm:"notNull"`
	User User
	PhotoID uint `json:"photo_id" gorm:"notNull"`
	Photo Photo
	Message string `json:"message" gorm:"type:text;notNull"`
	CreatedAt time.Time `json:"created_at" gorm:"notNull"`
	UpdatedAt time.Time `json:"updated_at" gorm:"notNull"`
}

type RequestComment struct{
	Message string `json:"message" valid:"required"`
	PhotoID uint `json:"photo_id,omitempty"`
}

type CommentIncludeUserPhoto struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	PhotoID uint `json:"photo_id"`
	Message string `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User struct {
		ID string `json:"id"`
		Email string `json:"email"`
		Username string `json:"username"`
	}
	Photo struct{
		ID string `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoUrl string `json:"photo_url"`
		UserID   uint	`json:"user_id"`
	}
}