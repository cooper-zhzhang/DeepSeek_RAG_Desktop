package model

import "gorm.io/gorm"

type DatasetModel struct {
	CollectName string `gorm:"unique"`
	gorm.Model
}

type FileModel struct {
	FilePath string
	Md5      string `gorm:"unique"`
	Status   int16  //状态 0:未处理 1:处理中 2:处理成功 3:处理失败

	gorm.Model
}

type DatasetRelaFileModel struct {
	DatasetId uint64
	FileId    uint64
	gorm.Model
}
