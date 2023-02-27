package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Id    int64 `gorm:"primaryKey"`
	Code  string
	Price uint
}
