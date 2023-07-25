package channelmessages

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

func (co *Controller) GetChannelMessage(c echo.Context) error {
	cmi, err := strconv.Atoi(c.Param("channel_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	cm, err := co.r.GetById(cmi)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, cm)
}

func (co *Controller) GetChannelMessages(c echo.Context) error {
	cms, err := co.r.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	if len(*cms) == 0 {
		return echo.NewHTTPError(http.StatusNoContent, utils.ErrorMessage{Message: ""})
	}

	return c.JSON(http.StatusOK, cms)
}

func (co *Controller) CreateChannelMessage(c echo.Context) error {
	cm := new(ChannelMessage)
	if err := c.Bind(cm); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	m, isEmpty := utils.ValidationEmptyValues[ChannelMessage](*cm)
	if isEmpty {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: m})
	}

	cmn, err := co.r.Create(*cm)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrorMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, cmn)
}
