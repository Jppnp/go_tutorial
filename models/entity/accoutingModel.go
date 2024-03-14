package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Accounting struct {
	Id        uint            `gorm:"primaryKey,autoIncrement" json:"id"`
	Amount    decimal.Decimal `gorm:"type:decimal;not null" json:"amount"`
	Balance   decimal.Decimal `gorm:"type:decimal;not null" json:"balance"`
	UserId    uint            `gorm:"not null" json:"userId"`
	User      User            `gorm:"foreignKey:UserId;references:ID" json:"-"`
	CreatedAt time.Time       `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time       `gorm:"autoCreateTime:false"`
}

func (Accounting) TableName() string {
	return "go_tutorial.accounting"
}
