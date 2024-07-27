package main

import (
	"app/POS/database"
	"app/POS/initializers"
	"app/POS/routes"
)

func init() {
	initializers.LoadEnvVars()
}

func main() {
	database.SetupDatabase()
	r := routes.SetupRouter()
	r.Run()
}
