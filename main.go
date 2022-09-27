package main

import (
	"structure-code-part2/config"
	"structure-code-part2/routes"

)


func main() {
	// initialize database
	config.InitDB()
	// auto migrate table and create table if not exist 
	config.InitMigrate()
	// initialize routes
	e := routes.New()
	// port
	e.Logger.Fatal(e.Start(":8000"))
}
