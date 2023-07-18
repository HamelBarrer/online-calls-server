package user

import (
	"net/http"
	"strconv"

	"github.com/HamelBarrer/online-calls-server/internal/utils"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	r Repository
}

func NewController(r Repository) *Controller {
	return &Controller{r}
}

func (co *Controller) GetUser(c echo.Context) error {
	ui, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	u, err := co.r.GetById(ui)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, u)
}

func (co *Controller) GetUsers(c echo.Context) error {
	us, err := co.r.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	if len(*us) == 0 {
		return echo.NewHTTPError(http.StatusNoContent, utils.ErrorMessage{Message: ""})
	}

	return c.JSON(http.StatusOK, us)
}

func (co *Controller) Create(c echo.Context) error {
	u := new(UserCreate)

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusNoContent, utils.ErrorMessage{Message: err.Error()})
	}

	m, isEmpty := utils.ValidationEmptyValues[UserCreate](*u)

	if isEmpty {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: m})
	}

	if u.Password != u.PasswordConfirm {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: "The password not equals"})
	}

	un, err := co.r.Create(*u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, un)
}
