package main

import (
	"github.com/labstack/echo"

	"log"

	"github.com/KitlerUA/freetoroom/db"
	"github.com/KitlerUA/freetoroom/ftrmid"
	"github.com/KitlerUA/freetoroom/handlers"
)

func main() {
	//connect to db
	if err := db.Connect(); err != nil {
		log.Fatalf("Cannot create db connections: %s", err)
	}
	defer db.Close()

	//create Echo
	e := echo.New()

	e.GET("/login", handlers.Login)
	e.GET("/logout", handlers.Logout)
	adminGroup := e.Group("/admin")
	adminGroup.Use(ftrmid.CheckCookie)
	adminGroup.POST("/room", handlers.AddRoom)
	adminGroup.PUT("/room", handlers.Update)
	adminGroup.DELETE("/room", handlers.Delete)
	adminGroup.GET("/room", handlers.GetAll)

	//start Server
	e.Start(":1323")
}
