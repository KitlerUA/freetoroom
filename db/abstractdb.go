package db

import (
	"github.com/KitlerUA/freetoroom/hotel"
)

//AbstractDB - interface with basic methods for database
type AbstractDB interface {
	AddRoom(room int, name string) error
	UpdateRoom(room int, name string) error
	DeleteRoom(room int) error
	GetAllRooms() ([]hotel.Room, error)
	CheckAccount(username, password string) (bool, error)
}
