package config

import (
	"Assigment2/structs"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/assigment2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected to Database")
	db.AutoMigrate(structs.Items{}, structs.Orders{})
	// db.Model(&structs.Items{}).AddForeignKey("o_id", "orders(Order_id)", "RESTRICT", "RESTRICT")

	db.Migrator().CreateConstraint(&structs.Orders{}, "Item")
	db.Migrator().CreateConstraint(&structs.Orders{}, "fk_orders_to_items")

	DB = db
}
