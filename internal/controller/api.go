package controller

import (
  "strconv"
  "os"
  "net/http"
  "flag"
  "io"
  "fmt"
)

import (
  "github.com/a-u-93/task/internal/aspect"
  "github.com/a-u-93/task/internal/view"
)

type API struct { *aspect.Aspect }

func (api *API) dayExchange(w http.ResponseWriter, r *http.Request) {
  year, err := strconv.Atoi(r.PathValue("year"))
  if err != nil || year < 1 { http.Error(w, "year specification is invalid ",
    404) }
  month, err := strconv.Atoi(r.PathValue("month"))
  if err != nil || month < 1 { http.Error(w, "month specification is invalid ",
    404) }
  day, err := strconv.Atoi(r.PathValue("day"))
  if err != nil || day < 1 { http.Error(w, "day specification is invalid ",
    404) }
  if year == 0 || month == 0 || day == 0 {
    http.Error(w, "date specification is invalid", 404)
  }
  renderedURL := fmt.Sprintf("%s?ondate=%d-%d-%d&periodicity=0",
    os.Getenv("UPSTREAM_API"), year, month, day)
  api.Info("rendered upstream API URL", "url", renderedURL)
  response, err := http.Get(renderedURL)
  if err != nil {
    api.Error("upstream API request", "status", "failed")
    http.Error(w, "Internal Server Error", 404)
  }
  body, err := io.ReadAll(response.Body)
  if err != nil {
    api.Error("upstream API response parsing", "status", "failed")
    http.Error(w, "Upstream API response parsing failed.", 404)
  }
  // api.Info("upstream API response", "body", string(body))
  view.DayExchangePersisted(api.Aspect, w, string(body), year, month, day)
}

func (api *API) historyExchange(w http.ResponseWriter, r *http.Request) {
  view.HistoryExchangeRendered(api.Aspect, w)
}

func (api *API) Listening() {
  address := flag.String("address", os.Getenv("MIDDLEWARE_ADDRESS"),
    "HTTP web address")
  mux := &http.ServeMux{}
  mux.HandleFunc("/{year}/{month}/{day}", api.dayExchange)
  mux.HandleFunc("/{$}", api.historyExchange)
  api.Info("Service should be started on 127.0.0.1:7777")
  err := http.ListenAndServe(*address, mux)
  if err != nil {
    api.Error("HTTP server had shuttered down", "message", err)
    os.Exit(1)
  }
}
