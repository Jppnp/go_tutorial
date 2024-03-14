package businessflow

import (
	"go/tutorial/config"
	"go/tutorial/models"
	"go/tutorial/models/entity"
	"go/tutorial/stage"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func CreateAccounting(tx *gorm.DB, user entity.User) (*entity.Accounting, error) {
	accounting := entity.Accounting{
		Amount:     decimal.New(100, 0),
		Balance:    decimal.New(100, 0),
		UserId:     user.ID,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		OutdatedBy: nil,
	}
	if err := stage.CreateAccounting(tx, &accounting); err != nil {
		return nil, err
	}
	return &accounting, nil
}

func PaymentProcess(tx *gorm.DB, user entity.User) (*models.AccountingStageModel, error) {
	oldAccounting := entity.Accounting{}
	if err := config.DB.First(&oldAccounting, "balance != ? AND outdated_by IS NULL", 0).Error; err != nil {
		return nil, err
	}
	accountStageModel := models.AccountingStageModel{
		OldAccouting: oldAccounting,
		NewAccounting: entity.Accounting{
			Amount:    decimal.New(100, 0),
			Balance:   decimal.New(0, 0),
			UserId:    user.ID,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}
	if err := stage.PaymentAccounting(tx, &accountStageModel); err != nil {
		return nil, err
	}
	return &accountStageModel, nil
}
