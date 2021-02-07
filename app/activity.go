package app

import (
  "fmt"
)

func GetAndIncrement(id int) (string, error) {
  db, err := ConnectToDatabase()
  if err != nil {
    fmt.Println(err)
    return fmt.Sprintf("ERROR: %s", err), err
  }
  defer CloseDatabaseConnection(db)
  obj, err := IncAndGet(1, db)
  if err != nil {
    fmt.Println(err)
    return fmt.Sprintf("ERROR: %s", err), err
  }
  result := fmt.Sprintf("Timestamp:\t %s \nValue:\t\t %d", obj.timestamp, obj.value)
  fmt.Println(result)
  return result, nil
}
