package main

import (
	"log"
	"net/http"
	"streaming-api/config/db"

	"github.com/labstack/echo/v4"
)

func main() {
	_, err := db.OpenConn()

	if err != nil {
		log.Fatalf("error connect database")
	}

	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.Start(":8080")
}
