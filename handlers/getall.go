package handlers

import (
	"strconv"

	"github.com/KitlerUA/freetoroom/db"
	"github.com/labstack/echo"
)

func GetAll(c echo.Context) error {
	records, err := db.GetAllRooms()
	if err != nil {
		return c.HTML(501, "Cannot get records: "+err.Error())
	}
	s := ""
	for i := range records {
		s += "Room: " + strconv.Itoa(records[i].Room) + "; Client: " + records[i].Name + "\n"
	}
	return c.HTML(ok, s)
}
