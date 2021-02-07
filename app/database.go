package app

import (
  "fmt"
  "log"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"

)

type timestampedInt struct {
  id int
  timestamp string
  value int
}

func ConnectToDatabase() (*sql.DB, error){
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/arbetsprov")
  if err != nil {
      return nil, err
  } else {
    fmt.Println("Established database connection")
    return db, nil
  }
}

func CloseDatabaseConnection(db *sql.DB) {
  fmt.Println("Closing database connection")
  db.Close()
}

func getTimestampedInt(id int, db *sql.DB) (timestampedInt, error){
  sqlQuery := fmt.Sprintf("SELECT timestamp, value FROM timestamped_int WHERE id = %d", id)

  var object timestampedInt
  results, err := db.Query(sqlQuery)
  if err != nil {
      return object, err
  }


  if !results.Next() {
    log.Fatalln("No Results from DB")
  }
  err = results.Scan(&object.timestamp, &object.value)
  object.id = id
  return object, err
}

func IncAndGet(id int, db *sql.DB) (timestampedInt, error) {
  _, err := db.Query("UPDATE timestamped_int SET timestamp = NOW(), value = value + 1 WHERE id = ?", id)
  var newValue timestampedInt
  if err != nil {
    return newValue, err
  }
  newValue, err = getTimestampedInt(id, db)
  return newValue, err
}
