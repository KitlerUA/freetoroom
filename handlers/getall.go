package handlers

import (
	"strconv"

	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) GetAll(c echo.Context) error {
	records, err := h.DB.GetAllRooms()
	if err != nil {
		return c.HTML(501, "Cannot get records: "+err.Error())
	}
	s := ""
	for i := range records {
		s += "Room: " + strconv.Itoa(records[i].Room) + "; Client: " + records[i].Name + "\n"
	}
	return c.HTML(http.StatusOK, s)
}
