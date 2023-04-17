package model

import "time"

type Comment struct {
	CommentID string `gorm:"primaryKey;type:varchar(255)"`
	Message   string `gorm:"not null;type:varchar(255);default:null"`
	UserID    string
	PhotoID   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//	Request
type CommentCreateRequest struct {
	Message string `json:"comment" valid:"required~insert your comment!"`
}

type CommentUpdateRequest struct {
	Message string `json:"comment" valid:"required~insert your comment!"`
}

//	Response
type CommentCreateResponse struct {
	CommentID string    `json:"comment_id"`
	Message   string    `json:"message"`
	UserID    string    `json:"user_id"`
	PhotoID   string    `json:"photo_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentUpdateResponse struct {
	CommentID string    `json:"comment_id"`
	Message   string    `json:"message"`
	UserID    string    `json:"user_id"`
	PhotoID   string    `json:"photo_id"`
	UpdatedAt time.Time `json:"update_at"`
}

type CommentResponse struct {
	CommentID string    `json:"comment_id"`
	Message   string    `json:"message"`
	UserID    string    `json:"user_id"`
	PhotoID   string    `json:"photo_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}
