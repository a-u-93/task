package view

import (
  "net/http"
  "encoding/json"
  "strconv"
)

import (
  "github.com/a-u-93/task/internal/aspect"
  "github.com/a-u-93/task/internal/model"
)

func DayExchangePersisted(
  api *aspect.Aspect, w http.ResponseWriter, jsonDocument string,
  year, month, day int,
) {
  w.Header().Add("Content-Type", "application/json")
  dayExchangeDocument := new(model.DayExchange)
  persisted := dayExchangeDocument.PersistencyCheck(api, year, month, day)
  api.Info("persistency check", "exists", strconv.FormatBool(persisted))
  if persisted {
    w.Write([]byte(jsonDocument))
    return
  }
  err := json.Unmarshal([]byte(jsonDocument), &dayExchangeDocument)
  if err != nil {
    api.Error("unmarshalling error", "message", err)
    w.Write([]byte(`{{"day currency exchange": "failed"}}`))
  }
  // api.Info("unmarshalled upstream API document", "document", dayExchangeDocument)
  err = dayExchangeDocument.Persisted(api)
  if err != nil {
    api.Error("day exchange data persistency", "status", "failed")
    w.Write([]byte(`{{"day currency exchange": "failed"}}`))
  }
  json.NewEncoder(w).Encode([]model.Currency(*dayExchangeDocument))
}

func HistoryExchangeRendered(api *aspect.Aspect, w http.ResponseWriter) {
  w.Header().Add("Content-Type", "application/json")
  var historyExchangeInstance model.HistoryExchange =
    make([]model.Currency, 0, 4096)
  err := historyExchangeInstance.Loaded(api)
  if err != nil {
    api.Error("currency history load error", "message", err)
    w.Write([]byte(`{{"history currency exchange fetch": "failed"}}`))
    // http.Error(w, "Internal Server Error", 500)
  }
  // api.Info("test fetched history", "list", historyExchangeInstance)
  // json.NewEncoder(w).Encode(historyExchangeInstance)
  // raw, err := json.Marshal(historyExchangeInstance)
  // if err != nil {
  //   api.Error("unmarshalling error", "message", err)
  //   http.Error(w, "Internal Server Error", 500)
  // }
  // json.NewEncoder(w).Encode([]model.Currency(historyExchangeInstance))
  data, _ := json.Marshal([]model.Currency(historyExchangeInstance))
  w.Write(data)
}
