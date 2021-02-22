package config

import (
	// "Assigment2/structs"

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
	db.AutoMigrate(structs.Items{}, structs.Orders{})

	fmt.Println("Connected to Database")
	// db.Model(&structs.Items{}).AddForeignKey("o_id", "orders(Order_id)", "RESTRICT", "RESTRICT")

	// db.Migrator().CreateConstraint(&structs.Orders{}, "Items")
	// db.Migrator().CreateConstraint(&structs.Orders{}, "fk_orders_to_items")

	// db.Migrator().HasConstraint(&structs.Orders{}, "Items")
	// db.Migrator().HasConstraint(&structs.Orders{}, "fk_orders_to_items")

	// db.Migrator().DropConstraint(&structs.Orders{}, "Items")
	// db.Migrator().DropConstraint(&structs.Orders{}, "fk_orders_to_items")

	DB = db
}
