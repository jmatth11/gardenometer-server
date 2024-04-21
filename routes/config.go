package routes

import (
	"database/sql"
	"gardenometer/actions"
	"gardenometer/email"

	"github.com/labstack/echo/v4"
)

func createConfig(conn *sql.DB, actionCache *actions.Queue, emailHandler *email.EmailClient) echo.HandlerFunc {
  return func(c echo.Context) error {

    return nil
  }
}

// TODO flesh out form values
//func configFormDataToModel(c echo.Context) models.Config {
//  conf := models.Config{}
//  name := c.FormValue("config_name")
//  wait := c.FormValue("config_wait")
//  moisture_pin := c.FormValue("moisture_pin")
//  temp_pin := c.FormValue("temp_pin")
//  lux_pin := c.FormValue("lux_pin")
//  cal_pin := c.FormValue("cal_pin")
//  err_pin := c.FormValue("err_pin")
//  good_pin := c.FormValue("good_pin")
//
//  return conf
//}
