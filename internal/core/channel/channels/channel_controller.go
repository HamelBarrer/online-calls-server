package channels

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

func (co *Controller) GetChannel(c echo.Context) error {
	ci, err := strconv.Atoi(c.Param("channel_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	ch, err := co.r.GetById(ci)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ch)
}

func (co *Controller) GetChannels(c echo.Context) error {
	cs, err := co.r.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	if len(*cs) == 0 {
		return echo.NewHTTPError(http.StatusNoContent, utils.ErrorMessage{Message: ""})
	}

	return c.JSON(http.StatusOK, cs)
}

func (co *Controller) CreateChannel(c echo.Context) error {
	ch := new(Channel)
	if err := c.Bind(ch); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	m, isEmpty := utils.ValidationEmptyValues[Channel](*ch)
	if isEmpty {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: m})
	}

	cha, err := co.r.Create(*ch)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, cha)
}
