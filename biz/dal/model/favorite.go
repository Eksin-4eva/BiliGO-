package model

import "time"

// Favorite 点赞/收藏表
// 索引建议：(user_id, video_id) 联合唯一索引防重复点赞，video_id 索引（统计点赞数）
type Favorite struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"                          json:"id"`
	UserID    int64     `gorm:"not null;uniqueIndex:idx_user_video"               json:"user_id"`
	VideoID   int64     `gorm:"not null;uniqueIndex:idx_user_video;index"         json:"video_id"`
	CreatedAt time.Time `gorm:"autoCreateTime"                                    json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"                                    json:"updated_at"`
}

func (Favorite) TableName() string { return "favorites" }
