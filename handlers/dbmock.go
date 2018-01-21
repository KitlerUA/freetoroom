package handlers

import (
	"fmt"
	"github.com/KitlerUA/freetoroom/hotel"
)

type DBMock struct {
	DB map[string]map[int]string
}

func (d *DBMock) AddRoom(room int, name string) error {
	if _, ok := d.DB["rooms"][room];ok{
		return fmt.Errorf("already exist")
	} else {
		d.DB["rooms"][room] = name
		return nil
	}
}

func (d *DBMock) UpdateRoom(room int, name string) error {
	if _, ok := d.DB["rooms"][room];!ok{
		return fmt.Errorf("not exist")
	} else {
		d.DB["rooms"][room] = name
		return nil
	}
}

func (d *DBMock) DeleteRoom(room int) error {
	if _, ok := d.DB["rooms"][room];!ok{
		return fmt.Errorf("not exist")
	} else {
		return nil
	}
}

func (d *DBMock) GetAllRooms() ([]hotel.Room, error){
	var res []hotel.Room
	for i := range d.DB["rooms"] {
		res = append(res, hotel.Room{Room: i,Name:d.DB["rooms"][i]})
	}
	if len(d.DB["rooms"]) > 0 {
		return res, nil
	} else {
		return res, fmt.Errorf("empty result")
	}
}

func (d *DBMock) CheckAccount(username, password string) (bool, error) {
	//account := &hotel.Account{}
	return true, nil
}