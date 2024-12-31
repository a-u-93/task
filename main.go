package main

import (
  // "log"
  // "net/http"
  "flag"
  "fmt"
  "log/slog"
  "os"
  "database/sql"
)

// import (
//   "github.com/a-u-93/task/internal/controller"
//   "github.com/a-u-93/task/internal/view"
//   "github.com/a-u-93/task/internal/model"
// )

import (
  _ "github.com/go-sql-driver/mysql"
)

type aspect struct {
  *slog.Logger
  *sql.DB
  environment
}

type environment []string

func main() {
  new(aspect).initialized()
  // controller.API(aspect)
  // controller.Cron(aspect)
}

func (a *aspect) initialized() {
  a.Logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
      Level: slog.LevelDebug, AddSource: true,
  }))

  a.DB = databaseConnection(a.Logger)

  a.environment = os.Environ()
}

func databaseConnection(logger *slog.Logger) *sql.DB {
  dsnDefault := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
    os.Getenv("MARIADB_USER"), os.Getenv("MARIADB_PASSWORD"),
      os.Getenv("MARIADB_ADDRESS"), os.Getenv("MARIADB_DATABASE"))
  dsn := flag.String("dsn", dsnDefault, "MariaDB data source")
  logger.Info("dsn", "rendered value", *dsn)
  flag.Parse()
  db, err := sql.Open("mysql", *dsn)
  if err != nil {
    message := "could not ti establish MariaDB connection"
    logger.Error("MariaDB connection", "status", "not established")
    panic(fmt.Sprintf("[Panic] %s: %s", message, err))
  }
  logger.Info("MariaDB connection", "status", "established")
  err = db.Ping()
  if err != nil {
    message := "could not ping database via already have established connection"
    logger.Error("MariaDB connectivity", "check", "failed")
    panic(fmt.Sprintf("[Panic] %s: %s", message, err))
  }
  logger.Error("MariaDB connectivity", "check", "ok")
  return db
}


  // address := flag.String("address", os.Getenv("MIDDLEWARE_ADDRESS"),
  //   "HTTP web address")
  // mux := &http.ServeMux{}
  // mux.HandleFunc("/{year}/{month}/{day}", api.DayExchange)
  // mux.HandleFunc("/{$}", api.HistoryExchange)
  // log.Print("Service should be started on 127.0.0.1:7777")
  // log.Fatal(http.ListenAndServe("127.0.0.1:7777", mux))
