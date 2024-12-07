package routes

import (
	"database/sql"
	"gardenometer/actions"
	"gardenometer/db"
	"gardenometer/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func upsertConfig(conn *sql.DB, actionCache *actions.Queue) echo.HandlerFunc {
  return func(c echo.Context) error {
    var configObj models.ConfigData
    err := c.Bind(&configObj)
    if err != nil {
      log.Println(err)
      return c.String(http.StatusBadRequest, "error with request")
    }
    existingConf, err := db.ReadConfigForDevice(conn, configObj.Name)
    if existingConf == nil && err == nil {
      err := db.InsertConfigForDevice(conn, &configObj)
      if err != nil {
        log.Println(err)
        return c.String(http.StatusInternalServerError, "error inserting config")
      }
    } else if err != nil {
      log.Println(err)
      return c.String(http.StatusInternalServerError, "error reading config")
    } else {
      err = db.UpdateConfigForDevice(conn, &configObj)
      if err != nil {
        log.Println(err)
        return c.String(http.StatusInternalServerError, "error updating config")
      }
    }
    actionCache.Push(actions.GenerateConfigAction(configObj.Name, configObj.ToConfig()))
    return c.String(http.StatusOK, "")
  }
}

func getConfig(conn *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    id := c.Param("id")
    if id == "" {
      return c.String(http.StatusBadRequest, "error: Must supply an ID")
    }
    conf, err := db.ReadConfigForDevice(conn, id)
    if conf == nil && err == nil {
      return c.JSON(
        http.StatusAccepted,
        models.ConfigData{
          Name: id,
          Wait: 1000,
          MoistureAir: 0,
          MoistureWater: 0,
        },
      )
    } else if err != nil {
      log.Println(err)
      return c.String(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, conf)
  }
}

