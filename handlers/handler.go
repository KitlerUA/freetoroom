package handlers

import (
	"github.com/labstack/echo"
	"github.com/KitlerUA/freetoroom/db"
)

//Handler - contains database field and methods for work in hotel
type Handler struct {
	DB db.AbstractDB
}