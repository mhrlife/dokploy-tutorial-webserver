package main

import (
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	echoServer := echo.New()

	echoServer.GET("/api/test", func(c echo.Context) error {
		welcomeTxt := os.Getenv("WELCOME_TEXT")

		if welcomeTxt == "" {
			welcomeTxt = "Hello world"
		}

		return c.String(200, welcomeTxt)
	})

	echoServer.Logger.Error(echoServer.Start(":8000"))
}
