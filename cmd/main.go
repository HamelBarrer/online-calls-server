package main

import (
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

	ur := user.Newrepository(db)
	uc := user.NewEchoRouter(e)
	uc.SetupRoute(ur)

	e.Logger.Fatal(e.Start(":3000"))
}
