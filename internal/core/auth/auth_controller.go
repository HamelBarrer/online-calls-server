package auth

import (
	"net/http"

	"github.com/HamelBarrer/online-calls-server/internal/utils"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	r Repository
}

func NewController(r Repository) *Controller {
	return &Controller{r}
}

func (co *Controller) Login(c echo.Context) error {
	au := new(Auth)
	if err := c.Bind(au); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	m, isEmpty := utils.ValidationEmptyValues[Auth](*au)
	if isEmpty {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: m})
	}

	u, err := co.r.GetUserByUsername(au.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	isEquals, err := utils.VerifyHash(au.Password, u.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	if !isEquals {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: "Username or password incorrect"})
	}

	t, err := utils.CreationToken(u.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user_id":  u.UserId,
		"username": u.Username,
		"token":    t,
	})
}
