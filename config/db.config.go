package config

import (
	"fmt"
	"go/tutorial/exception"
	"log"
	"math/rand"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	dbCfg := Cfg.DB_CONNECTION
	dsn := dbCfg.CONNECTION_STRING
	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      loggerDb,
	})
	exception.PanicLog(err)
	DB = db
	instance, _ := DB.DB()
	instance.SetMaxOpenConns(int(dbCfg.MAX_POOL))
	instance.SetMaxIdleConns(int(dbCfg.IDLE_POOL))
	instance.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(dbCfg.LIFE_TIME))) * time.Millisecond)

}
func DBTransaction() *gorm.DB {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	return tx
}

func DBCommit(tx *gorm.DB) error {
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("เกิดข้อผิดพลาดในการบันทึกข้อมูล error:%w", err)
	}
	return nil
}
