package entity

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserName  string    `gorm:"type:varchar(255);not null" json:"userName"`
	TypeUser  uint      `gorm:"type:int;not null" json:"typeUser"`
	CreatedAt time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time `gorm:"autoCreateTime:false"`
}
