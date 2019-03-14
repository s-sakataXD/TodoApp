package model

import "time"

// Todo はDBにマッピングされている構造体
type Todo struct {
	ID        int `gorm:"primary_key"`
	Content   string
	CreatedAt time.Time
}
