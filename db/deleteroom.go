package db

import (
	"github.com/KitlerUA/freetoroom/hotel"
	"fmt"
)

//DeleteRoom - delete given room from database
func (d *GormDB) DeleteRoom(room int) error {
	if d.db.Unscoped().Where("room = ?", room).Delete(&hotel.Room{}).RowsAffected == 0 {
		return fmt.Errorf("record not found")
	}
	return nil
}
