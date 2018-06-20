package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Id        uint `gorm:"primary_key"`
	AuthorID  int  `gorm:"index"`
	BookID    int  `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Post) TableName() string {
	return "post"
}
