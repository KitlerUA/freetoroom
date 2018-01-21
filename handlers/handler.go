package handlers

import "github.com/labstack/echo"

type AbstractHandler interface {
	AddRoom(c echo.Context) error
	Delete(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	GetAll(c echo.Context) error
	Update(c echo.Context) error
}

type Handler struct {

}