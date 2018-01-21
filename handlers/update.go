package handlers

import (
	"strconv"

	"github.com/KitlerUA/freetoroom/db"
	"github.com/labstack/echo"
	"net/http"
)

func Update(c echo.Context) error {
	room := c.FormValue("room")
	if room == "" {
		return c.HTML(http.StatusBadRequest, "parameter 'room' not found")
	}
	var roomInt int64
	var err error
	roomInt, err = strconv.ParseInt(room, 10, 64)
	if err != nil {
		return c.HTML(http.StatusBadRequest, "bad parameter 'room': "+err.Error())
	}
	client := c.FormValue("client")
	if client == "" {
		return c.HTML(http.StatusBadRequest, "parameter 'client' not found")
	}

	if err = db.UpdateRoom(int(roomInt), client); err != nil {
		return c.HTML(501, "Cannot update room: "+err.Error())
	}

	return c.HTML(http.StatusOK, "Room "+strconv.Itoa(int(roomInt))+" was changed")
}
