package db

import "github.com/KitlerUA/freetoroom/hotel"

//AddRoom - insert Room into database
func (d *GormDB) AddRoom(room int, name string) error {
	return d.db.Create(&hotel.Room{Room: room, Name: name}).Error
}