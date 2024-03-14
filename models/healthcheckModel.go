package models

type Healthcheck struct {
	ID      uint   `gorm:"primaryKey;not null" json:"id"`
	Message string `gorm:"type:varchar(255);not null" json:"message"`
}

func (Healthcheck) TableName() string {
	return "go_tutorial.healthcheck"
}
