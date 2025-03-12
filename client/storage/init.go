package storage

import (
	"dp_client/storage/model"
	"dp_client/storage/query"
	"log"
	"time"

	"gorm.io/gorm/logger"

	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DeepSeekDB *gorm.DB
)

func init() {
	// 从配置文件中读取值
	dsn := viper.GetString("db.dsn")

	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // io.Writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // 设置日志级别为 Silent
			Colorful:      false,         // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.Agent{}, &model.Conversation{}, &model.Message{})
	if err != nil {
		panic("failed to migrate database")
	}
	DeepSeekDB = db

	query.SetDefault(DeepSeekDB)
}
