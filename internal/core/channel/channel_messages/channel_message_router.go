package channelmessages

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

	er.e.GET("/api/v1/channel_messages/:channel_id", c.GetChannelMessage)
	er.e.GET("/api/v1/channel_messages", c.GetChannelMessages)
	er.e.POST("/api/v1/channel_messages", c.CreateChannelMessage)
}
