package model

import (
	query "github.com/aide-cloud/gorm-normalize"
)

type FileType int8

type File struct {
	query.BaseModel
	UserID   uint     `gorm:"column:user_id;type:int(10) unsigned;not null;default:0;comment:用户ID" json:"user_id"`
	Name     string   `gorm:"column:name;type:varchar(20);not null;default:'';comment:文件名" json:"name"`
	Url      string   `gorm:"column:url;type:varchar(255);not null;default:'';comment:文件地址" json:"url"`
	FileType FileType `gorm:"column:file_type;type:tinyint(1);not null;default:0;comment:文件类型" json:"file_type"`
	Ext      string   `gorm:"column:ext;type:varchar(10);not null;default:'';comment:文件后缀" json:"ext"`
	Size     int64    `gorm:"column:size;type:bigint(20) unsigned;not null;default:0;comment:文件大小" json:"size"`
}

const (
	fileTableName = "files"
)

const (
	// FileTypeImage 图片
	FileTypeImage FileType = iota + 1
	// FileTypeVideo 视频
	FileTypeVideo
	// FileTypeAudio 音频
	FileTypeAudio
	// FileTypeDocument 文档
	FileTypeDocument
	// FileTypeOther 其他
	FileTypeOther
)

func (File) TableName() string {
	return fileTableName
}
