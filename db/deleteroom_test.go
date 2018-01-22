package db

import (
	"testing"
	"github.com/KitlerUA/freetoroom/hotel"
)

func TestGormDB_DeleteRoom(t *testing.T) {
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
		{"normal case", 1666, "client 1666", true},
		{"not exist", 666, "client 666", false},
	}
	for _, test := range testCases {
		testLocal := test
		t.Run(testLocal.name, func(t *testing.T){
			t.Parallel()
			if err := database.DeleteRoom(testLocal.room); err == nil{
				if testLocal.expected {
					r := database.db.Where("room = ?", testLocal.room).First(&hotel.Room{}).Value.(*hotel.Room)
					if r.Room != 0 || r.Name != "" {
						t.Errorf("Expected %d: %s , got %d: %s", 0, "-", r.Room, r.Name)
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