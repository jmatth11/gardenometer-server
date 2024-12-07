package routes

import (
	"database/sql"
	"gardenometer/actions"
	"gardenometer/db"
	"gardenometer/email"
	"gardenometer/models"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func upsertConfig(conn *sql.DB, actionCache *actions.Queue, emailHandler *email.EmailClient) echo.HandlerFunc {
  return func(c echo.Context) error {
    body := c.Request().Body
    conf := new(models.ConfigRequest)
    _, err := io.Copy(conf, body)
    if err != nil {
      log.Println(err)
      return c.String(http.StatusInternalServerError, "error handling request")
    }
    configObj := conf.ParseConfig()
    existingConf, err := db.ReadConfigForDevice(conn, configObj.Name)
    if existingConf == nil && err == nil {
      err := db.InsertConfigForDevice(conn, configObj)
      if err != nil {
        log.Println(err)
        return c.String(http.StatusInternalServerError, "error inserting config")
      }
    }
    if err != nil {
      log.Println(err)
      return c.String(http.StatusInternalServerError, "error reading config")
    }
    err = db.UpdateConfigForDevice(conn, configObj)
    if err != nil {
      log.Println(err)
      return c.String(http.StatusInternalServerError, "error updating config")
    }
    return nil
  }
}

