package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"middleware_loft/internal/app/endpoint"
	"middleware_loft/internal/app/mw"
	"middleware_loft/internal/app/service"
)

type App struct {
	endpoint *endpoint.Endpoint
	service  *service.Service
	echo     *echo.Echo
}

func New() (*App, error) {
	app := &App{}
	app.service = service.New()
	app.endpoint = endpoint.New(app.service)

	app.echo = echo.New()

	// use middleware
	app.echo.Use(mw.RoleCheck)

	app.echo.GET("/health", app.endpoint.Status)

	return app, nil
}

// TODO: add port
func (app *App) Run() error {
	fmt.Println("Server running >>>")

	err := app.echo.Start(":8081")

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
