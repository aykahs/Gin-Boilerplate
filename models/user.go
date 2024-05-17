package models

import "time"

type User struct {
	BasicModel

	Name        string    `json:"name"`                                 // Name
	Username    string    `json:"username" gorm:"uniqueIndex;not null"` // Username
	Password    string    `json:"password"`                             // Password
	LastLoginAt time.Time `json:"last_login_at" gorm:"default:null"`    // Last login time
}
