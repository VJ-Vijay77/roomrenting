package main

import (
	"github.com/VJ-Vijay77/r4room/pkg/database"
	"github.com/VJ-Vijay77/r4room/pkg/gingonic"
	"github.com/VJ-Vijay77/r4room/pkg/routes"
)



func main() {
	Engine := gingonic.GinInit()
	database.GetDb()
	routes.InitRoutes(Engine)
	Engine.Run()
}
