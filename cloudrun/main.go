package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mmorito/spanner/src/middleware"
	"github.com/mmorito/spanner/src/routes"
	"github.com/mmorito/spanner/src/utilities/logger"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// router
	routes.Router(e)

	e.Use(middleware.SetupLogger)

	// init
	initHandler(e)
}

func initHandler(e *echo.Echo) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), logger.Init(e)))
}
