package channels

import "github.com/labstack/echo/v4"

type Router interface {
	SetupRoute(Repository)
}

type EchoRouter struct {
	e *echo.Echo
}

func NewEchoRouter(e *echo.Echo) Router {
	return &EchoRouter{e}
}

func (er *EchoRouter) SetupRoute(r Repository) {
	c := NewController(r)

	er.e.GET("/api/v1/channels/:channel_id", c.GetChannel)
	er.e.GET("/api/v1/channels", c.GetChannels)
	er.e.POST("/api/v1/channels", c.CreateChannel)
}
