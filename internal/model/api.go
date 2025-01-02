package model

import (
  "time"
  "strconv"
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

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
  var vanillaTime *time.Time = new(time.Time)
  unquotedString, err := strconv.Unquote(string(data))
  if err != nil { panic(err) }
  *vanillaTime, err = time.Parse("2006-01-02T00:00:00", unquotedString)
  if err != nil { panic(err) }
  *ct = CustomTime(*vanillaTime)
  return nil
}

func (ct CustomTime) String() string {
  return time.Time(ct).String()
}

func (de DayExchange) Persisted(api *aspect.Aspect) error {
  api.Info("time debug", "time", time.Time(de[0].Date).String())
  sqlQuery := "replace into currency values(?, ?, ?, ?, ?, ?)"
  tx, err := api.Begin()
  if err != nil {
    api.Error("transaction start", "status", err)
    return err
  }
  defer tx.Rollback()
  for _, c := range de {
    _, err = tx.Exec(sqlQuery, c.CurId, time.Time(c.Date), c.CurAbbreviation, c.CurScale, c.CurName,
      c.CurOfficialRate)
    if err != nil {
      api.Error("currency persistency", "status", err)
      return err
    }
  }
  tx.Commit()
  return nil
}
