package main

import (
  "github.com/a-u-93/task/internal/controller"
  "github.com/a-u-93/task/internal/aspect"
)

func main() {
  api := new(controller.API)
  api.Aspect = aspect.Version["v0.0.1"]
  api.Listening()
}
