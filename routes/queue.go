package routes

import (
	"gardenometer/actions"
	"gardenometer/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func createQueue(actionCache *actions.Queue) echo.HandlerFunc {
  return func(c echo.Context) error {
    qr := new(models.QueueRequest)
    if err := c.Bind(qr); err != nil {
      return c.String(http.StatusBadRequest, "invalid payload")
    }
    switch(qr.Type) {
      case models.ActionCalibrate:{
        actionCache.Push(actions.GenerateCalibrationAction(qr.Name))
        break;
      }
      case models.ActionConfig: {
        if qr.Configuration == nil {
          return c.String(http.StatusBadRequest, "configuration required")
        }
        actionCache.Push(actions.GenerateConfigAction(qr.Name, *qr.Configuration))
        break;
      }
      case models.ActionCode: {
        if qr.Code == nil {
          return c.String(http.StatusBadRequest, "code required")
        }
        actionCache.Push(actions.GenerateCodeResponse(qr.Name, *qr.Code))
        break;
      }
    }
    return c.String(http.StatusOK, "")
  }
}

