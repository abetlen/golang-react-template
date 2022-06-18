package main

import (
	"embed"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

//go:embed all:ui/build
var build embed.FS

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	e := echo.New()
	e.GET("/api/healthz", func(c echo.Context) error {
		status := struct {
			Status string `json:"status"`
		}{
			Status: "OK",
		}
		if err := c.Bind(status); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, status)
	})
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "ui/build",
		Index:      "index.html",
		Browse:     false,
		HTML5:      true,
		Filesystem: http.FS(build),
	}))
	e.Logger.Fatal(e.Start(host + ":" + port))
}
