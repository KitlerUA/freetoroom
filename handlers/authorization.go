package handlers

import (
	"net/http"

	"time"

	"github.com/labstack/echo"
	"crypto/sha1"
	"encoding/hex"
)

//Login - if user exists, create cookie with token
func (h *Handler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	sh := sha1.Sum([]byte(password))
	ok, err := h.DB.CheckAccount(username, hex.EncodeToString(sh[:]))
	if err!=nil{
		return c.HTML(http.StatusUnauthorized, err.Error())
	}
	if ok {
		cookie := &http.Cookie{}
		cookie.Name = "sessionID"
		cookie.Value = "static_token"
		cookie.Secure = false
		cookie.Expires = time.Now().Add(12 * time.Hour)

		c.SetCookie(cookie)

		return c.String(http.StatusOK, "You were logged in")
	}
	return c.HTML(http.StatusUnauthorized, "Wrong credential")
}

//Logout - delete cookie
func (h *Handler) Logout(c echo.Context) error {
	cookie := &http.Cookie{}
	cookie.Name = "sessionID"
	cookie.Value = "deleted"
	cookie.Secure = false
	cookie.Expires = time.Now()

	c.SetCookie(cookie)

	return c.String(http.StatusOK, "You were logged out")
}