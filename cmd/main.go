package main

import (
	"github.com/HamelBarrer/online-calls-server/internal/core/auth"
	channelstates "github.com/HamelBarrer/online-calls-server/internal/core/channel/channel_states"
	"github.com/HamelBarrer/online-calls-server/internal/core/channel/channels"
	"github.com/HamelBarrer/online-calls-server/internal/core/user/user"
	"github.com/HamelBarrer/online-calls-server/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := storage.NewPostgres()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ar := auth.Newrepository(db)
	ac := auth.NewEchoRouter(e)
	ac.SetupRoute(ar)

	ur := user.Newrepository(db)
	uc := user.NewEchoRouter(e)
	uc.SetupRoute(ur)

	csr := channelstates.Newrepository(db)
	csc := channelstates.NewEchoRouter(e)
	csc.SetupRoute(csr)

	cr := channels.Newrepository(db)
	cc := channels.NewEchoRouter(e)
	cc.SetupRoute(cr)

	e.Logger.Fatal(e.Start(":3000"))
}
