package models

import (
	"time"
)

// User represents a user in the system.
type User struct {
    ID        uint      `gorm:"primary_key" json:"id"`
    FirstName string    `gorm:"type:varchar(100)" json:"first_name"`
    LastName  string    `gorm:"type:varchar(100)" json:"last_name"`
    Email     string    `gorm:"type:varchar(100);uniqueIndex" json:"email"`
    Password  string    `gorm:"type:varchar(100)" json:"password"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}