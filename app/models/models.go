package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Age       uint      `json:"age"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}

type Photo struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}

type Comment struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID"`
	PhotoID   uint      `json:"photo_id"`
	Photo     Photo     `gorm:"foreignKey:PhotoID"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}

type SocialMedia struct {
	gorm.Model
	ID              uint      `json:"id" gorm:"primary_key"`
	Name            string    `json:"name"`
	SocialMedialUrl string    `json:"social_media_url"`
	UserID          uint      `json:"user_id"`
	User            User      `gorm:"foreignKey:UserID"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoCreateTime,autoUpdateTime"`
}
