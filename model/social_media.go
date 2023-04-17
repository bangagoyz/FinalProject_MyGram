package model

import (
	"time"
)

type SocialMedia struct {
	SocialID       string `gorm:"primaryKey;type:varchar(255)"`
	Name           string `gorm:"not null;type:varchar(255);default:null"`
	SocialMediaUrl string `gorm:"not null;type:varchar(255);default:null"`
	UserID         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Request
type SocialMediaCreateRequest struct {
	Name           string `json:"name" valid:"required~Social Media Name required!"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Social Media URL Required!"`
}

type SocialMediaUpdateRequest struct {
	Name           string `json:"name" valid:"required~Social Media Name required!"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Social Media URL Required!"`
}

// Response
type SocialMediaCreateResponse struct {
	SocialID       string    `json:"social_id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         string    `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type SocialMediaUpdateResponse struct {
	SocialID       string    `json:"social_id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         string    `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaResponse struct {
	SocialID       string    `json:"social_id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         string    `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
