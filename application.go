package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/n3m/covid-dashboard-server/infrastructure/database"
	"github.com/n3m/covid-dashboard-server/infrastructure/router"
	"github.com/thedevsaddam/gojsonq/v2"
)

func main() {
	app := router.Init()

	DB := database.Init("./COVID/export_dataframe.json")

	/* MIDDLEWARE SETUP PHASE 1 */
	app.Use(DBMiddleware(DB))

	/* SERVER HANDLER SETUP P2 */
	router.SetRoutes(app)

	/* INIT */
	go app.Listen(":5000")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Printf("> Closing Server...")
	app.Shutdown()
}

// DBMiddleware ...
func DBMiddleware(DB *gojsonq.JSONQ) fiber.Handler {
	injector := func(c *fiber.Ctx) error {
		c.Locals("db", DB)
		return c.Next()
	}

	return injector
}
