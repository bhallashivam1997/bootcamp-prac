//main.go
package main

import (
	"bootcamp-prac/day-4/Config"
	"bootcamp-prac/day-4/Models"
	"bootcamp-prac/day-4/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	//Config.DB.AutoMigrate().CreateTable(&Models.Product{})
	Config.DB.AutoMigrate(Models.Product{},Models.Order{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}