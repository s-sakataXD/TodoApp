package Model

import "time"

type Todo struct {
	Id        int `gorm:"primary_key"`
	Content   string
	CreatedAt time.Time
}
