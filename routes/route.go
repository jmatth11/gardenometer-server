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

func getCalibrate(conn *sql.DB, actionCache *actions.Queue) echo.HandlerFunc {
  return func (c echo.Context) error {
    id := c.Param("id")
    if id == "" {
      return c.Render(http.StatusBadRequest, "error", "Must supply an ID")
    }
    req, err := db.ReadRegistration(conn, id)
    if err != nil {
      log.Println(err)
      return c.Render(http.StatusInternalServerError, "error", err)
    }
    if req == nil {
      return c.Render(
        http.StatusNotFound,
        "toast",
        ErrorToast("Device is not registered"),
      )
    }
    actionCache.Push(actions.GenerateCalibrationAction(req.Name))
    return c.Render(
      http.StatusOK, "toast", SuccessToast("Calibration request queued."))
  }
}

func getFlipIsActive(conn *sql.DB) echo.HandlerFunc {
  return func (c echo.Context) error {
    id := c.Param("id")
    if id == "" {
      return c.Render(http.StatusBadRequest, "error", "Must supply an ID")
    }
    req, err := db.ReadRegistration(conn, id)
    if err != nil {
      log.Println(err)
      return c.Render(http.StatusInternalServerError, "error", err)
    }
    if req == nil {
      return c.Render(
        http.StatusNotFound, "toast", ErrorToast("Device is not registered."))
    }
    req.IsActive = !req.IsActive
    err = db.UpdateRegistrationIsActive(conn, req.Name, req.IsActive)
    if err != nil {
      log.Println(err)
      return c.Render(http.StatusInternalServerError, "error", err)
    }
    c.Render(
      http.StatusOK, "toast", SuccessToast("Active State Changed"))
    return c.Render(
      http.StatusOK, "active_update", req)
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
  e.GET("/calibrate/:id", getCalibrate(conn, actionCache))
  e.GET("/change-active/:id", getFlipIsActive(conn))
}

