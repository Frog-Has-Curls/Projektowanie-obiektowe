package main
import (
    "fmt"
    "sync"
    "database/sql"

    _"github.com/mattn/go-sqlite3"
)
var (
    baza_danych *sql.DB
)

var lock = &sync.Mutex{}

type single struct {}

var singleInst *single

func getInstance() *single {
    if singleInst == nil {
        lock.Lock()
        defer lock.Unlock()
	 if singleInst == nil {
             fmt.Println("Creating single instance now.")
	     singleInst = &single{}
    } else {
	fmt.Println("Single instance already created.")
    }


    return singleInstance
}
 
func open_db() {
    database, err :=sql.Open("sqlite3", "baza_danych.db")
    if err !=nil{
panic(err)
}

    fmt.Println(database)
    fmt.Println("Connection opened")
}

func close_db() error {
fmt.Println("Database closed")
    return baza_danych.Close()
}



func(c echo.Context) error {
  rodzaj := c.QueryParam("rodzaj")
  return c.String(http.StatusOK, rodzaj)
})

