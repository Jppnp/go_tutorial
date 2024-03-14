package stage

import (
	"fmt"
	"go/tutorial/models"
	"go/tutorial/models/entity"

	"gorm.io/gorm"
)

func CreateAccounting(tx *gorm.DB, accounting *entity.Accounting) error {
	if err := tx.Create(accounting).Error; err != nil {
		return err
	}
	return nil
}

func PaymentAccounting(tx *gorm.DB, accountStageModel *models.AccountingStageModel) error {
	newAccouting := accountStageModel.NewAccounting
	if err := tx.Create(&newAccouting).Error; err != nil {
		return fmt.Errorf("create payment" + err.Error())
	}
	oldAccounting := accountStageModel.OldAccouting
	oldAccounting.OutdatedBy = &newAccouting.Id
	if err := tx.Save(oldAccounting).Error; err != nil {
		return fmt.Errorf("update old accounting" + err.Error())
	}
	return nil
}
