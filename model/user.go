package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm: "size:255;not null;unique" jason:"username"`
	Password string `gorm: "size:255;not null;" json:"-"`
	Entries  []Entry
}
