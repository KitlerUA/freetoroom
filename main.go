package main

import (
	"github.com/labstack/echo"

	"log"

	"github.com/KitlerUA/freetoroom/db"
	"github.com/KitlerUA/freetoroom/ftrmid"
	"github.com/KitlerUA/freetoroom/handlers"
	"fmt"
	"github.com/KitlerUA/freetoroom/config"
)

func main() {
	//connect to db
	dataBase := db.GormDB{}
	if err := dataBase.Connect("db/hotel"); err != nil {
		log.Fatalf("Cannot create db connections: %s", err)
	}
	defer dataBase.Close()

	h := handlers.Handler{DB: &dataBase}

	//create Echo
	e := echo.New()

	e.GET("/login", h.Login)
	e.GET("/logout", h.Logout)
	adminGroup := e.Group("/admin")
	adminGroup.Use(ftrmid.CheckCookie)
	adminGroup.POST("/room", h.AddRoom)
	adminGroup.PUT("/room", h.Update)
	adminGroup.DELETE("/room", h.Delete)
	adminGroup.GET("/room", h.GetAll)

	//start Server
	e.Start(fmt.Sprintf(":%s", config.Get().Port))
}
