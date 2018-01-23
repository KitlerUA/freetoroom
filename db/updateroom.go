package db

import (
	"github.com/KitlerUA/freetoroom/hotel"
	"fmt"
)

//UpdateRoom - change 'room'`s owner to 'name'
func (d *GormDB) UpdateRoom(room int, name string) error {
	if d.db.Model(&hotel.Room{}).Where("room = ?", room).Update("name", name).RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}
