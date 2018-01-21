package hotel

import "github.com/jinzhu/gorm"

type Room struct {
	gorm.Model

	Room int    `gorm:"not null;unique"`
	Name string `gorm:"not null"`
}
