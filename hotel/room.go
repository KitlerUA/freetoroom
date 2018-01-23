package hotel

import "github.com/jinzhu/gorm"

//Room - represent reserved room in hotel in database
type Room struct {
	gorm.Model

	Room int    `gorm:"not null;unique"`
	Name string `gorm:"not null"`
}
