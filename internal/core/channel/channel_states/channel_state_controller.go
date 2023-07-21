package channelstates

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

func (co *Controller) GetChannelState(c echo.Context) error {
	ci, err := strconv.Atoi(c.Param("channel_state_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	cs, err := co.r.GetById(ci)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, cs)
}

func (co *Controller) GetChannelStates(c echo.Context) error {
	cs, err := co.r.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	if len(*cs) == 0 {
		return echo.NewHTTPError(http.StatusNoContent, utils.ErrorMessage{Message: ""})

	}

	return c.JSON(http.StatusOK, cs)
}

func (co *Controller) ChannelStateCreate(c echo.Context) error {
	cs := new(ChannelState)
	if err := c.Bind(cs); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	m, isEmpty := utils.ValidationEmptyValues[ChannelState](*cs)
	if isEmpty {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: m})
	}

	state, err := co.r.Create(*cs)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, state)
}
