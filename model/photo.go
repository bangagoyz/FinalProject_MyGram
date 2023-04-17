package model

import (
	"time"
)

type Photo struct {
	PhotoID   string `gorm:"primaryKey;type:varchar(255)"`
	Title     string `gorm:"not null;type:varchar(255);default:null"`
	PhotoUrl  string `gorm:"not null;type:varchar(255);default:null"`
	UserID    string
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Request
type PhotoRequest struct {
	Title    string `json:"title" valid:"required~Photo Title is Required"`
	PhotoUrl string `json:"photo_url" valid:"required~Photo URL is required"`
}

// Response
type PhotoCreateResponse struct {
	PhotoID   string    `json:"photo_id"`
	Title     string    `json:"title"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoUpdateResponse struct {
	PhotoID   string    `json:"photo_id"`
	Title     string    `json:"title"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoResponse struct {
	PhotoID   string    `json:"photo_id"`
	Title     string    `json:"title"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	Comments  []Comment `json:"comments"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoAllResponse struct {
	PhotoID   string    `json:"photo_id"`
	Title     string    `json:"title"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
