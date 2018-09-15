package main

import (
  "fmt"
  "github.com/jinzhu/gorm"                            // ORM package for Go
  _ "github.com/jinzhu/gorm/dialects/sqlite"          // for SQLite. Only imports functions so that ORM can use. Hence the '_'
)

// This is the structure for your database. Very similar to how SQLAlchemy works with Flask.
type Person struct {
  ID uint `json:"id"`
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
}

func main() {
   db, _ := gorm.Open("sqlite3", "./gorm.db")             // Creates an SQLite database, stores it in the file.
   defer db.Close()                                       // defer basically tells it to execute at the end of the main function's scope
   db.AutoMigrate(&Person{})                              // Creates database based on the Person structure as schema
   p1 := Person{FirstName: "John", LastName: "Doe"}
   // p2 := Person{FirstName: "Jane", LastName: "Smith"}     // Example Person objects being created
   db.Create(&p1)                                         // Creates an entry in the db with the object p1
   var p3 Person                                          // identify a Person type for us to store the results in
   db.First(&p3)                                          // Find the first record in the Database and store it in p3
   fmt.Println(p3.FirstName)
   fmt.Println(p3.LastName)                               // print out our record from the database
}
