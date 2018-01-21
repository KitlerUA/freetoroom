package handlers

import "github.com/labstack/echo"

func Anything(c echo.Context) error{
	return c.HTML(400, c.Path())
}
