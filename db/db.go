package db

import (
	"github.com/jinzhu/gorm"

	"time"

	"github.com/KitlerUA/freetoroom/hotel"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

var db *gorm.DB

func Get(d *GormDB) *gorm.DB {
	return d.db
}

func (d *GormDB) Connect(address string) error {
	var err error
	d.db, err = gorm.Open("sqlite3", address)
	d.db.AutoMigrate(&hotel.Room{})
	d.db.AutoMigrate(&hotel.Account{})
	return err
}

func (d *GormDB) AddRoom(room int, name string) error {
	return d.db.Create(&hotel.Room{Room: room, Name: name}).Error
}

func (d *GormDB) UpdateRoom(room int, name string) error {
	if d.db.Model(&hotel.Room{}).Where("room = ?", room).Update("name", name).RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}

func (d *GormDB) DeleteRoom(room int) error {
	if d.db.Unscoped().Where("room = ?", room).Delete(&hotel.Room{}).RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}

func (d *GormDB) GetAllRooms() ([]hotel.Room, error) {
	var res []hotel.Room
	rows, err := d.db.Raw("select * from rooms").Rows()
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

func (d *GormDB) CheckAccount(username, password string) (bool, error) {
	account := &hotel.Account{}
	if err := d.db.Where("username = ? and password = ?", username, password).First(account).Error;err!=nil{
		return false, err
	}
	return true, nil
}

func (d *GormDB) Close() error {
	return d.db.Close()
}
