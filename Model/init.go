package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" //DBドライバを直接使用すべきではないため
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres",
		"user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Todo{})
}
