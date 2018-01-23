package db

import (
	"github.com/KitlerUA/freetoroom/hotel"
	"time"
)

//GetAllRooms - return slice of booked rooms
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
