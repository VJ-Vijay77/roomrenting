package main

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/gingonic"
	"github.com/VJ-Vijay77/r4room/pkg/routes"
)



func main() {
	//getting gin engine
	Engine := gingonic.GinInit()

	//connecting database
	database.GetDb()

	//activating routes
	routes.InitRoutes(Engine)
	
	//starting the port
	Engine.Run()
}
