package handlers

import (
	"fmt"
	"github.com/KitlerUA/freetoroom/hotel"
)

//DBMock - mock for database
type DBMock struct {
	DB map[string]map[int]string
	DBAccounts map[string]map[string]string
}

//AddRoom - mock for AddRoom method
func (d *DBMock) AddRoom(room int, name string) error {
	if _, ok := d.DB["rooms"][room];ok{
		return fmt.Errorf("already exist")
	}
	d.DB["rooms"][room] = name
	return nil
}

//UpdateRoom - mock for UpdateRoom method
func (d *DBMock) UpdateRoom(room int, name string) error {
	if _, ok := d.DB["rooms"][room];!ok{
		return fmt.Errorf("not exist")
	}
	d.DB["rooms"][room] = name
	return nil
}

//DeleteRoom - mock for DeleteRoom method
func (d *DBMock) DeleteRoom(room int) error {
	if _, ok := d.DB["rooms"][room];!ok{
		return fmt.Errorf("not exist")
	}
	return nil
}

//GetAllRooms - mock for GetAllRooms method
func (d *DBMock) GetAllRooms() ([]hotel.Room, error){
	var res []hotel.Room
	for i := range d.DB["rooms"] {
		res = append(res, hotel.Room{Room: i,Name:d.DB["rooms"][i]})
	}
	return res, nil
}

//CheckAccount - mock for CheckAccount method
func (d *DBMock) CheckAccount(username, password string) (bool, error) {
	if username == "kitler" && password == "e5e9fa1ba31ecd1ae84f75caaa474f3a663f05f4" {
		return true, nil
	}
	return false, fmt.Errorf("bad credential")
}