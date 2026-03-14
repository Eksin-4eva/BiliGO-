package model

import "time"

// Comment 评论表
// 索引建议：video_id 索引（查评论列表），parent_id 索引（查子评论）
type Comment struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"   json:"id"`
	VideoID   int64     `gorm:"not null;index"             json:"video_id"`
	UserID    int64     `gorm:"not null"                   json:"user_id"`
	Content   string    `gorm:"type:text;not null"         json:"content"`
	ParentID  int64     `gorm:"default:0;index"            json:"parent_id"` // 0 表示顶级评论
	CreatedAt time.Time `gorm:"autoCreateTime"             json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"             json:"updated_at"`
}

func (Comment) TableName() string { return "comments" }
