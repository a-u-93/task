package aspect

import (
  "log/slog"
  "flag"
  "fmt"
  "os"
  "database/sql"
)

import _ "github.com/go-sql-driver/mysql"

// We define an application as a set of cross-cutting-concerns, named as
// aspects here.

// Problem here is polymorphic Aspect design - depending on version, aspect
// changes it's structure.
type Aspect struct {
  *slog.Logger
  *sql.DB
  // Environment
}

var Version map[string]*Aspect = map[string]*Aspect{
  "v0.0.1": devVersion(),
  // "v2.0.0": v2.0.0(),
}

// To-Do: replace with `Viper` and get YAML config files support, `etcd`
// changes listening mode and `go` template engine config rendering.
// type Environment map[string]string


func devVersion() *Aspect {
  a := new(Aspect)

  a.Logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
      Level: slog.LevelDebug, AddSource: true,
  }))

  a.DB = databaseConnection(a.Logger)

  // a.Environment = &Environment{}.readed()

  return a
}

// `Viper` configuration values should be promoted here.
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
