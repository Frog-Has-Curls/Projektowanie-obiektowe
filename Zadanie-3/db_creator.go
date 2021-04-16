package main

import (
         "database/sql"
          "fmt"
          "strconv"
      _ "github.com/mattn/go-sqlite3"
    )

   func main() {
     database, _ := sql.Open("sqlite3", "./baza_danych.db")
     statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS baza_danych (id INTEGER PRIMARY KEY, rodzaj TEXT, gatunek TEXT)")
     statement.Exec()
     statement, _ = database.Prepare("INSERT INTO baza_danych (rodzaj, gatunek) VALUES (?, ?)")
     statement.Exec("Bacillus", "subtilis")
     rows, _ := database.Query("SELECT id, rodzaj, gatunek FROM baza_danych")

     var id int
     var rodzaj string
     var gatunek string
     for rows.Next() {
         rows.Scan(&id, &rodzaj, &gatunek)
          fmt.Println(strconv.Itoa(id) + ": " + rodzaj + " " + gatunek)
}
}
