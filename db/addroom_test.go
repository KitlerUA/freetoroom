package db

import (
	"testing"
	"github.com/KitlerUA/freetoroom/hotel"
)

func TestGormDB_AddRoom(t *testing.T) {
	database := GormDB{}
	if err := database.Connect("testdb");err!=nil{
		t.Errorf("Cannot conect to database: %s",err)
	}

	database.db.Unscoped().Model(&hotel.Room{}).Delete(&hotel.Room{})
	database.db.Unscoped().Model(&hotel.Account{}).Delete(&hotel.Account{})

	database.db.Exec(insertRoom, 1666, "client 1666")

	testCases := []struct {
		name string
		room int
		client string
		expected bool
	}{
		{"normal case", 666, "client 666", true},
		{"already exist", 1666, "client 1666", false},
	}
	for _, test := range testCases {
		testLocal := test
		t.Run(testLocal.name, func(t *testing.T){
			t.Parallel()
			if err := database.AddRoom(testLocal.room, testLocal.client); err == nil{
				if testLocal.expected {
					r := database.db.Where("room = ?", testLocal.room).First(&hotel.Room{}).Value.(*hotel.Room)
					if r.Room != testLocal.room || r.Name != testLocal.client {
						t.Errorf("Expected %d: %s , got %d: %s", testLocal.room, testLocal.client, r.Room, r.Name)
					}
				} else {
					t.Errorf("Expected error , got success", )
				}
			} else {
				if testLocal.expected {
					t.Errorf("Expected success , got error: %s", err )
				}
			}
		})
	}
}
