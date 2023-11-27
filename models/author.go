package models

import "time"

type Author struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Gender    string    `gorm:"type:char(1)" json:"gender"`
	Email     string    `gorm:"type:varchar(100)" json:"email"`
	Age       string    `gorm:"type:integer" json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
