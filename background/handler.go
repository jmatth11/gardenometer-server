package background

import (
	"database/sql"
	"gardenometer/db"
	"gardenometer/email"
	"log"
	"strings"
	"time"
)

func Start(exit chan bool, conn *sql.DB, emailClient *email.EmailClient) {
  pulseCheck := time.NewTimer(time.Minute * 30)
  run := true
  for {
    select {
    case <-exit:
      run = false;
    case <-pulseCheck.C:
      checkDeviceActivity(conn, emailClient)
      pulseCheck = time.NewTimer(time.Minute * 30)
    }
    if !run {
      break
    }
  }
  exit <- true
}

func checkDeviceActivity(conn *sql.DB, emailClient *email.EmailClient) {
  activities, err := db.ReadLatestMetricForEachName(conn)
  if err != nil {
    err := emailClient.SendMail(err.Error())
    if err != nil {
      log.Println(err)
    }
  }
  now := time.Now()
  timeThreshold := time.Minute * 10
  sb := strings.Builder{}
  shouldEmail := false
  for _, e := range activities {
    // skip inactive devices
    if !e.IsActive {
      continue
    }
    diff := now.Sub(e.UpdatedAt)
    if diff > timeThreshold {
      shouldEmail = true
      sb.WriteString(e.Name)
      sb.WriteString(" has been offline for more than 10 minutes\n")
      sb.WriteString("Last activity: ")
      sb.WriteString(e.UpdatedAt.String())
      sb.WriteString("\n")
      // adding this for debugging
      sb.WriteString("now: ")
      sb.WriteString(now)
      sb.WriteString("\n")
      sb.WriteString("Diff: ")
      sb.WriteString(diff)
      sb.WriteString("\n")
      sb.WriteString("timeThreshold: ")
      sb.WriteString(timeThreshold)
      sb.WriteString("\n")
    }
  }
  if shouldEmail {
    err := emailClient.SendMail(sb.String())
    if err != nil {
      log.Println(err)
    }
  }
}
