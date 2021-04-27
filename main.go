package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	id      int `gorm:"primary_key"`
	Name    string
	Address string
}

func main() {

	db, err := gorm.Open(sqlite.Open("NASA_STATS.db"), &gorm.Config{})
	if err != nil {
		log.Println("Connection to Database has failed to open")
	}

	log.Println("Connection to database established.")

	db.AutoMigrate(&User{})

	db.Create(&User{}) // Figure out how to convert json data into a struct that can be used to update sqlite.

	req, err := http.Get("https://api.nasa.gov/insight_weather/?api_key=0urOUYPFipLzQ9ZF40IvOrFZ429ql8ns7XlwJV0O&feedtype=json&ver=1.0")
	if err != nil {
		log.Fatal(err)
	}
	read, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", read)
}
