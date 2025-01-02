package controller

import (
  "time"
  "fmt"
  "os"
  "net/http"
)

import (
  "github.com/a-u-93/task/internal/aspect"
)

type Cron struct { *aspect.Aspect }

func (c *Cron) Started() {
  ticker := time.NewTicker(24 * time.Hour)
  for {
    select {
    case <- ticker.C:
      c.CurrentDayExchange()
    }
  }
}

func (c *Cron) CurrentDayExchange() {
  c.Info("cron job started")
  now := time.Now()
  upstreamAPI := fmt.Sprintf("%s?ondate=%d-%d-%d&periodicity=0",
    os.Getenv("UPSTREAM_API"), now.Year(), now.Month(), now.Day())
  _, err := http.Get(upstreamAPI)
  if err != nil {
    c.Error("cron upstream API data fetch", "error", err)
    return
  }
  c.Info("cron job ended")
}
