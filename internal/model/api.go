package model

import (
  "fmt"
  "time"
  "strconv"
  "encoding/json"
)

import (
  "github.com/a-u-93/task/internal/aspect"
)

type CustomTime time.Time

type Currency struct {
  CurId int `json:"Cur_ID"`
  Date CustomTime `json:"Date"`
  CurAbbreviation string `json:"Cur_Abbreviation"`
  CurScale int `json:"Cur_Scale"`
  CurName string `json:"Cur_Name"`
  CurOfficialRate float32 `json:"Cur_OfficialRate"`
}

type DayExchange []Currency

type HistoryExchange []Currency

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
  var vanillaTime *time.Time = new(time.Time)
  unquotedString, err := strconv.Unquote(string(data))
  if err != nil { panic(err) }
  *vanillaTime, err = time.Parse("2006-01-02T00:00:00", unquotedString)
  if err != nil { panic(err) }
  *ct = CustomTime(*vanillaTime)
  return nil
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
  return json.Marshal(time.Time(*ct))
}

func (ct CustomTime) String() string {
  return time.Time(ct).String()
}

func (de DayExchange) PersistencyCheck(api *aspect.Aspect, year, month, day int) bool {
  var count int
  var sqlQueryLayout string =
    "select count(`date`) as `count` from `currency` where `date` = '%d-%d-%d';"
  var sqlQuery string = fmt.Sprintf(sqlQueryLayout, year, month, day)
  tx, err := api.Begin()
  if err != nil {
    api.Error("transaction creation", "failed", err)
    return false
  }
  defer tx.Rollback()
  err = tx.QueryRow(sqlQuery).Scan(&count)
  if err != nil {
    api.Error("row scan", "failed", err)
    return false
  }
  err = tx.Commit()
  if err != nil {
    api.Error("transaction commit", "failed", err)
    return false
  }
  api.Info("day exchange persistency check", "count", count)
  if count == 0 {
    return false
  } else {
    return true
  }
}

func (de DayExchange) Persisted(api *aspect.Aspect) error {
  // api.Info("time debug", "time", time.Time(de[0].Date).String())
  sqlQuery := "replace into currency values(?, ?, ?, ?, ?, ?)"
  tx, err := api.Begin()
  if err != nil {
    api.Error("day currency persistency transaction start", "status - failed",
      err)
    return err
  }
  defer tx.Rollback()
  for _, c := range de {
    _, err = tx.Exec(sqlQuery, c.CurId, time.Time(c.Date), c.CurAbbreviation,
      c.CurScale, c.CurName, c.CurOfficialRate)
    if err != nil {
      api.Error("currency persistency", "status", err)
      return err
    }
  }
  tx.Commit()
  return nil
}

func (he *HistoryExchange) Loaded(api *aspect.Aspect) error {
  sqlQuery := "select `cur_id`, `date`, `cur_abbreviation`, `cur_scale`, `cur_name`, `cur_officialrate` from `currency`"
  tx, err := api.Begin()
  if err != nil {
    api.Error("currency history transaction start", "status - failed", err)
    return err
  }
  defer tx.Rollback()
  rows, err := api.DB.Query(sqlQuery)
  if err != nil {
    api.Error("currency history rows fetch error", "status - failed", err)
    return err
  }
  defer rows.Close()
  for rows.Next() {
    var c Currency
    var t time.Time
    err = rows.Scan(&c.CurId, &t, &c.CurAbbreviation, &c.CurScale,
      &c.CurName, &c.CurOfficialRate)
    if err != nil {
      api.Error("currency record scan err", "failed", err)
      return err
    }
    c.Date = CustomTime(t)
    *he = append(*he, c)
  }
  err = rows.Err()
  if err != nil {
    api.Error("currency history rows reading error", "status - failed", err)
    return err
  }
  err = tx.Commit()
  if err != nil {
    api.Error("currency history transaction commit error", "status - failed",
      err)
    return err
  }
  return nil
}
