package main

import (
	"fmt"

	"github.com/twisty/clean-arch/database"
	"github.com/twisty/clean-arch/models"
	"github.com/twisty/clean-arch/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	databaseInfo := database.SetUpDb()
	database.DB, err = gorm.Open(mysql.Open(databaseInfo.ReadConfigDBFile()), &gorm.Config{})
	if err != nil {
		fmt.Println("Connect database error " + err.Error())
	}

	database.DB.AutoMigrate(&models.Product{}, &models.Order{}, &models.OrderDetail{}, &models.User{}, &models.Branch{},
		&models.Stock{})
	closeConnectionDB, _ := database.DB.DB()
	defer fmt.Println("Close Connection DB !!!")
	defer closeConnectionDB.Close()
	server := routes.SetUpRoutes()
	server.Run(":9999")

}
