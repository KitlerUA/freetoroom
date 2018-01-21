package ftrmid

import (
	"net/http"

	"github.com/labstack/echo"
)

func CheckCookie(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			return c.HTML(http.StatusBadGateway, "cookie error: "+err.Error())
		}
		if cookie.Value == "static_token" {
			return h(c)
		}
		return c.HTML(http.StatusUnauthorized, "Please, login and try again")
	}
}
