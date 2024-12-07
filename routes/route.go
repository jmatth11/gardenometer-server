package routes

import (
	"database/sql"
	"gardenometer/actions"
	"gardenometer/db"
	"gardenometer/email"
	"gardenometer/helpers"
	"gardenometer/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func getHome() echo.HandlerFunc {
  return func (c echo.Context) error {
    log.Println("calling from: ", c.Request().Header.Get("User-Agent"))
    return c.Render(http.StatusOK, "index", "")
  }
}

func getRegistrationList(conn *sql.DB) echo.HandlerFunc {
  return func (c echo.Context) error {
    registrationList, err := db.ReadLatestMetricForEachName(conn)
    if err != nil {
      log.Println(err)
      return c.Render(http.StatusInternalServerError, "error", err)
    }
    return c.Render(http.StatusOK, "registration_list", registrationList)
  }
}

func getConfigTab(conn *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    req, err := db.ReadAllRegistration(conn)
    if err != nil {
      log.Println(err)
      return c.Render(http.StatusInternalServerError, "error", err)
    }
    names := make([]string, 0, len(req))
    for _, v := range req {
      names = append(names, v.Name)
    }
    ct := models.ConfigTab{
      Devices: names,
    }
    return c.Render(http.StatusOK, "config_tab", ct)
  }
}

func ping(c echo.Context) error {
  log.Println("calling from: ", c.Request().Header.Get("User-Agent"))
  return c.String(http.StatusOK, "gardenometer")
}

func Setup(e *echo.Echo, conn *sql.DB, actionCache *actions.Queue, emailHandler *email.EmailClient) {
  e.Pre(middleware.RemoveTrailingSlash())
  e.Use(middleware.Recover())
  e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
      rate.Limit(20),
  )))

  e.Static("/css", "css")

  helpers.NewTemplateRenderer(e)

  e.GET("/", getHome())
  e.GET("/registration_list", getRegistrationList(conn))
  e.GET("/config_tab", getConfigTab(conn))
  e.GET("/status", ping)
  e.POST("/status", createStatus(conn, actionCache, emailHandler))
  e.POST("/queue", createQueue(actionCache))
  e.POST("/alert", createAlert(conn))
  e.POST("/config", createConfig(conn, actionCache, emailHandler))
}

