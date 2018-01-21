package db

import "github.com/jinzhu/gorm"
import (
	"sync"

	"time"

	"github.com/KitlerUA/freetoroom/hotel"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func Get(d *GormDB) *gorm.DB {
	return d.db
}

func (d *GormDB) Connect() error {
	var err error
	d.db, err = gorm.Open("sqlite3", "db/hotel")
	d.db.AutoMigrate(&hotel.Room{})
	d.db.AutoMigrate(&hotel.Account{})
	return err
}

func AddRoom(room int, name string) error {
	return db.Create(&hotel.Room{Room: room, Name: name}).Error
}

func UpdateRoom(room int, name string) error {
	return db.Model(&hotel.Room{}).Where("room = ?", room).Update("name", name).Error
}

func DeleteRoom(room int) error {
	return db.Unscoped().Where("room = ?", room).Delete(&hotel.Room{}).Error
}

func GetAllRooms() ([]hotel.Room, error) {
	var res []hotel.Room
	rows, err := db.Raw("select * from rooms").Rows()
	if err != nil {
		return res, err
	}
	for rows.Next() {
		//var room hotel.Room
		var id, room int
		var ct, ut, dt *time.Time
		var name string
		rows.Scan(&id, &ct, &ut, &dt, &room, &name)

		res = append(res, hotel.Room{Room: room, Name: name})
	}
	return res, nil
}

func CheckAccount(username, password string) (bool, error) {
	account := &hotel.Account{}
	if err := db.Where("username = ? and password = ?", username, password).First(account).Error;err!=nil{
		return false, err
	}
	return true, nil
}

func Close() error {
	return db.Close()
}
