// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameVideo = "video"

// Video mapped from table <video>
type Video struct {
	ID            int64                 `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`             // id主键
	AuthorID      int64                 `gorm:"column:author_id;type:bigint;not null" json:"author_id"`                    // 作者id
	Title         string                `gorm:"column:title;type:varchar(128);not null" json:"title"`                      // 视频标题
	PlayURL       string                `gorm:"column:play_url;type:varchar(255);not null" json:"play_url"`                // 视频播放地址
	CoverURL      string                `gorm:"column:cover_url;type:varchar(255);not null" json:"cover_url"`              // 视频封面地址
	FavoriteCount int64                 `gorm:"column:favorite_count;type:bigint;not null" json:"favorite_count"`          // 点赞数
	CommentCount  int64                 `gorm:"column:comment_count;type:bigint;not null" json:"comment_count"`            // 评论数
	CreateTime    int64                 `gorm:"column:create_time;type:bigint unsigned;autoCreateTime" json:"create_time"` // 创建时间
	DeletedAt     soft_delete.DeletedAt `gorm:"column:deleted_at;type:int;softDelete:milli" json:"deleted_at"`             // 软删除
}

// TableName Video's table name
func (*Video) TableName() string {
	return TableNameVideo
}
