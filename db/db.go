package db

import (
	"github.com/jinzhu/gorm"

	"github.com/KitlerUA/freetoroom/hotel"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//GormDB - implement AbstractDB using sqlite with gorm
type GormDB struct {
	db *gorm.DB
}

//Connect - create connection to database with given address
func (d *GormDB) Connect(address string) error {
	var err error
	d.db, err = gorm.Open("sqlite3", address)
	d.db.AutoMigrate(&hotel.Room{})
	d.db.AutoMigrate(&hotel.Account{})
	return err
}

//Close - close connection to database
func (d *GormDB) Close() error {
	return d.db.Close()
}
