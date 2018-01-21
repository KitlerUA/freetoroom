package handlers

import (
	"strconv"

	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) Delete(c echo.Context) error {
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

	if err = h.DB.DeleteRoom(int(roomInt)); err != nil {
		return c.HTML(501, "Cannot delete room: "+err.Error())
	}
	return c.HTML(http.StatusOK, "Room "+strconv.Itoa(int(roomInt))+" was deleted")
}
