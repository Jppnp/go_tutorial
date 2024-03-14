package stage

import (
	"go/tutorial/models/entity"

	"gorm.io/gorm"
)

func CreateUserStage(tDb *gorm.DB, user *entity.User) error {
	if err := tDb.Create(user).Error; err != nil {
		return err
	}
	return nil
}
