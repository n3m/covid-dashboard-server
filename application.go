package main

import (
	"log"

	"github.com/n3m/covid-dashboard-server/router"
)

func main() {
	app := router.Init("./data.json")

	defer log.Fatal(app.Listen(":3000"))
}
