package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/USACE/shared-api/middleware"
	"github.com/USACE/shared-api/static"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORS, middleware.GZIP)

	// Public Routes
	public := e.Group("")

	// Health Check
	public.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"status": "healthy"})
	})

	public.GET("/offices", func(c echo.Context) error {
		return c.JSONBlob(http.StatusOK, static.Offices)
	})

	log.Print("starting server")
	log.Fatal(http.ListenAndServe(":80", e))
}
