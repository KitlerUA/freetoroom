package hotel

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model

	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}
