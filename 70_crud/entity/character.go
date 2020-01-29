package entity

import "github.com/jinzhu/gorm"

type Character struct {
	gorm.Model
	Name string
}
