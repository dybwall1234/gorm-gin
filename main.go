package main

import (
	"fmt"
	"github.com/dybwall1234/gorm-gin/Config"
	"github.com/dybwall1234/gorm-gin/Models"
	"github.com/dybwall1234/gorm-gin/Routers"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql",
		Config.Dburl)

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Book{})

	r := Routers.SetupRouter()
	// running
	r.Run()
}
