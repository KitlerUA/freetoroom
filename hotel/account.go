package hotel

import "github.com/jinzhu/gorm"

//Account - represent user (admin user) in database
type Account struct {
	gorm.Model

	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}
