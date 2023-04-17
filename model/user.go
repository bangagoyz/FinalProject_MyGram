package model

import "time"

type User struct {
	ID           string `gorm:"primaryKey;type:varchar(255)"`
	Username     string `gorm:"unique;not null;type:varchar(255);default:null"`
	Email        string `gorm:"unique;not null;type:varchar(255);default:null"`
	Password     string `gorm:"not null;type:varchar(255)"`
	Age          int    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	SocialMedias []SocialMedia
	Photos       []Photo
	Comments     []Comment
}

// Request
type UserRegisterRequest struct {
	Username string `json:"username" valid:"required~Username is required"`
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required~Password is required,minstringlength(6)~Password atleast 6 characters"`
	Age      int    `json:"age" valid:"required~Age is required,range(8|99)~minimum age to register is 8"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required~Username is required"`
	Password string `json:"password" validate:"required~Password is required"`
}

// Response
type UserRegisterResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
