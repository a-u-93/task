package view

import (
  "net/http"
  "encoding/json"
)

import (
  "github.com/a-u-93/task/internal/aspect"
  "github.com/a-u-93/task/internal/model"
)

func DayExchangePersisted(api *aspect.Aspect, w http.ResponseWriter, jsonDocument string) {
  dayExchangeDocument := new(model.DayExchange)
  err := json.Unmarshal([]byte(jsonDocument), &dayExchangeDocument)
  if err != nil { api.Error("unmarshalling error", "messgae", err) }
  // api.Info("unmarshalled upstream API document", "document", dayExchangeDocument)
  err = dayExchangeDocument.Persisted(api)
  if err != nil {
    api.Error("day exchange data persistency", "status", "failed")
  }
}
