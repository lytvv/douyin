// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFollow = "follow"

// Follow mapped from table <follow>
type Follow struct {
	ID        int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`                         // id主键
	UserID    int64     `gorm:"column:user_id;type:bigint;not null" json:"user_id"`                                    // 用户id
	FollowID  int64     `gorm:"column:follow_id;type:bigint;not null" json:"follow_id"`                                // 关注用户id
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
}

// TableName Follow's table name
func (*Follow) TableName() string {
	return TableNameFollow
}
