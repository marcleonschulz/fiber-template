package database

import (
	"fmt"
	"github.com/marcleonschulz/fiber-template/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connectDb

func Init() error {
	conf := config.Config.Db
	dsn := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable",
		conf.Host,
		conf.User,
		conf.Password,
		conf.Name,
		conf.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		//PrepareStmt: true,
	})
	if err != nil {
		return err
	}
	//db.Logger = logger.Default.LogMode(logger.Info)
	err = db.AutoMigrate()

	DB = db
	if err != nil {
		return err
	}
	return nil
}

// HealthCheck give a health check for the database
func HealthCheck() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
