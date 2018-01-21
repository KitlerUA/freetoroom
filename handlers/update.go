package handlers

import (
	"strconv"

	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) Update(c echo.Context) error {
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

	if err = h.DB.UpdateRoom(int(roomInt), client); err != nil {
		return c.HTML(http.StatusInternalServerError, "Cannot update room: "+err.Error())
	}

	return c.HTML(http.StatusOK, "Room "+strconv.Itoa(int(roomInt))+" was changed")
}
