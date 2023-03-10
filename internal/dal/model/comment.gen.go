// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameComment = "comment"

// Comment mapped from table <comment>
type Comment struct {
	ID            int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`                         // id主键
	UserID        int64     `gorm:"column:user_id;type:bigint;not null" json:"user_id"`                                    // 用户id
	VideoID       int64     `gorm:"column:video_id;type:bigint;not null" json:"video_id"`                                  // 视频id
	FavoriteCount int64     `gorm:"column:favorite_count;type:bigint;not null" json:"favorite_count"`                      // 点赞数
	Content       string    `gorm:"column:content;type:varchar(255);not null" json:"content"`                              // 评论内容
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
