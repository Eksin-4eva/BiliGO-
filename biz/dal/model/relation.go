package model

import "time"

// Relation 关注关系表
// status: 1=关注, 0=取消关注
// 索引建议：(user_id, to_user_id) 联合唯一索引，to_user_id 索引（查粉丝列表）
type Relation struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"                            json:"id"`
	UserID    int64     `gorm:"not null;uniqueIndex:idx_relation"                   json:"user_id"`
	ToUserID  int64     `gorm:"not null;uniqueIndex:idx_relation;index"             json:"to_user_id"`
	Status    int8      `gorm:"default:1"                                           json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime"                                      json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"                                      json:"updated_at"`
}

func (Relation) TableName() string { return "relations" }
