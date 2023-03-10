// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID              int64                 `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`                         // id主键
	Name            string                `gorm:"column:name;type:varchar(128);not null" json:"name"`                                    // 用户名
	Password        string                `gorm:"column:password;type:varchar(128);not null" json:"password"`                            // 密码
	Avatar          string                `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`                                // 头像地址
	FollowCount     int64                 `gorm:"column:follow_count;type:bigint;not null" json:"follow_count"`                          // 关注总数
	FollowerCount   int64                 `gorm:"column:follower_count;type:bigint;not null" json:"follower_count"`                      // 粉丝总数
	TotalFavorited  int64                 `gorm:"column:total_favorited;type:bigint;not null" json:"total_favorited"`                    // 总点赞数
	Signature       string                `gorm:"column:signature;type:varchar(255);not null" json:"signature"`                          // 个性签名
	BackgroundImage string                `gorm:"column:background_image;type:varchar(255);not null" json:"background_image"`            // 背景图片
	WorkCount       int64                 `gorm:"column:work_count;type:bigint;not null" json:"work_count"`                              // 作品总数
	FavoriteCount   int64                 `gorm:"column:favorite_count;type:bigint;not null" json:"favorite_count"`                      // 点赞总数
	CreatedAt       time.Time             `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt       time.Time             `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
	DeletedAt       soft_delete.DeletedAt `gorm:"column:deleted_at;type:int;softDelete:milli" json:"deleted_at"`                         // 软删除
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
