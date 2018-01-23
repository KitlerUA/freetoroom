package db

import (
	"testing"
	"github.com/KitlerUA/freetoroom/hotel"
)

func TestGormDB_GetAllRooms_Normal(t *testing.T) {
	database := GormDB{}
	if err := database.Connect("testdb");err!=nil{
		t.Errorf("Cannot conect to database: %s",err)
	}

	database.db.Unscoped().Model(&hotel.Room{}).Delete(&hotel.Room{})
	database.db.Unscoped().Model(&hotel.Account{}).Delete(&hotel.Account{})

	database.db.Exec(insertRoom, 1666, "client 1666")
	database.db.Exec(insertRoom, 2666, "client 2666")
	database.db.Exec(insertRoom, 3666, "client 3666")

	testCases := []struct {
		name string
		expected []hotel.Room
	}{
		{"normal case", []hotel.Room{{Room:1666, Name:"client 1666"},{Room:2666, Name: "client 2666"}, {Room: 3666, Name:"client 3666"} }},
	}
	for _, test := range testCases {
		testLocal := test
		t.Run(testLocal.name, func(t *testing.T){
			t.Parallel()
			if rooms, err := database.GetAllRooms(); err == nil && len(testLocal.expected) == len(rooms){
				for i := range rooms {
					if rooms[i].Room != testLocal.expected[i].Room ||
						rooms[i].Name != testLocal.expected[i].Name {
						t.Errorf("Expected %s: %s , got %s: %s", testLocal.expected[i].Room, testLocal.expected[i].Name, rooms[i].Room, rooms[i].Name)
					}
				}
			} else {
				t.Errorf("Expected %v , got %v", testLocal.expected, rooms)
			}
		})
	}
}

func TestGormDB_GetAllRooms_Empty(t *testing.T) {
	database := GormDB{}
	if err := database.Connect("testdb");err!=nil{
		t.Errorf("Cannot conect to database: %s",err)
	}

	database.db.Unscoped().Model(&hotel.Room{}).Delete(&hotel.Room{})
	database.db.Unscoped().Model(&hotel.Account{}).Delete(&hotel.Account{})

	testCases := []struct {
		name string
		expected []hotel.Room
	}{
		{"empty database", []hotel.Room{ }},
	}
	for _, test := range testCases {
		testLocal := test
		t.Run(testLocal.name, func(t *testing.T){
			t.Parallel()
			if rooms, err := database.GetAllRooms(); err != nil || rooms != nil {
				t.Errorf("Expected nil-slice , got %v with error %s",rooms, err)
			}
		})
	}
}