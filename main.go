package main

import (
	"gardenometer/actions"
	"gardenometer/background"
	"gardenometer/db"
	"gardenometer/email"
	"gardenometer/routes"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
  env, err := godotenv.Read(".env")
  if err != nil {
    panic(err)
  }
  outputLog, err := os.Create("output.log")
  if err != nil {
    panic(err)
  }
  defer outputLog.Close()
  log.SetOutput(io.MultiWriter(os.Stdout, outputLog))
  e := echo.New()
  db, err := db.Setup(env["DB_USER"], env["DB"], env["DB_PW"])
  if err != nil {
    panic(err)
  }
  emailClient := email.NewEmailClient(env["EMAIL"], env["EMAIL_PW"])
  emailClient.To = []string{env["EMAIL_TO"]}
  var exit chan bool
  go background.Start(exit, db, emailClient)
  activeCache := actions.NewQueue()
  routes.Setup(e, db, activeCache, emailClient);
  err = e.Start(":8000")
  if err != nil {
    log.Println(err)
  }
  exit <- true
  <-exit
}
