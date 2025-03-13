package storage

import (
	"dp_client/global"
	"dp_client/storage/model"
	"dp_client/storage/query"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	// 从配置文件中读取值
	dsn := global.DBViper.GetString("db.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.Agent{}, &model.Conversation{}, &model.Message{}, &model.Agent{})
	if err != nil {
		panic("failed to migrate database")
	}
	query.SetDefault(db)
}

// TODO storage 层全部做成单例
