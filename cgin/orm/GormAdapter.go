package orm

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormAdapter struct {
	*gorm.DB
}

func NewGormAdapter() *GormAdapter {
	dsn := "root:123456@tcp(192.168.226.129:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dbHelper, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, _ := dbHelper.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &GormAdapter{DB: dbHelper}
}

func (this *GormAdapter) Name() string {
	return "GormAdapter"
}
