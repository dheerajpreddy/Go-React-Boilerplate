package main

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Person struct {
  ID uint `json:"id"`
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
}

func main() {
   db, _ := gorm.Open("sqlite3", "./gorm.db")
   defer db.Close()
   db.AutoMigrate(&Person{})
   p1 := Person{FirstName: "John", LastName: "Doe"}
   p2 := Person{FirstName: "Jane", LastName: "Smith"}
   db.Create(&p1)
   var p3 Person // identify a Person type for us to store the results in
   db.First(&p3) // Find the first record in the Database and store it in p3
   fmt.Println(p1.FirstName)
   fmt.Println(p2.LastName)
   fmt.Println(p3.LastName) // print out our record from the database
}
