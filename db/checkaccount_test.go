package db

import (
	"testing"
	"github.com/KitlerUA/freetoroom/hotel"
)

func TestGormDB_CheckAccount(t *testing.T) {
	database := GormDB{}
	if err := database.Connect("testdb");err!=nil{
		t.Errorf("Cannot conect to database: %s",err)
	}

	database.db.Unscoped().Model(&hotel.Room{}).Delete(&hotel.Room{})
	database.db.Unscoped().Model(&hotel.Account{}).Delete(&hotel.Account{})

	database.db.Exec(insertAccounts, "kitler", "e5e9fa1ba31ecd1ae84f75caaa474f3a663f05f4")

	testCases := []struct {
		name string
		username string
		password string
		expected bool
	}{
		{"normal case", "kitler", "e5e9fa1ba31ecd1ae84f75caaa474f3a663f05f4", true},
		{"not exist", "not kitler", "e5e9fa1ba31ecd1ae84f75caaa474f3a663f05f4", false},
		{"wrong password", "kitler", "E5e9fa1ba31ecd1ae84f75caaa474f3a663f05f4", false},
	}
	for _, test := range testCases {
		testLocal := test
		t.Run(testLocal.name, func(t *testing.T){
			t.Parallel()
			if ok, _ := database.CheckAccount(testLocal.username, testLocal.password); ok != testLocal.expected{
				t.Errorf("Expected %v , got %v", testLocal.expected, ok)
			}
		})
	}
}