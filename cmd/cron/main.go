package main

import (
  "github.com/a-u-93/task/internal/controller"
  "github.com/a-u-93/task/internal/aspect"
)

func main() {
  cron := new(controller.Cron)
  cron.Aspect = aspect.Version["v0.0.1"]
  cron.Started()
}
