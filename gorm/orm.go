package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	//Create
	db.Create(&Product{Code: "D42", Price: 200})

	//Read
	var product Product
	db.First(&product, 1) //fetch with primary key
	db.First(&product, "code = ?", "D42")

	fmt.Println(product.Code)

	//Update
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	db.First(&product, 1) //fetch with primary key
	fmt.Println(product.Price)
	// Delete - delete product
	db.Delete(&product, 1)
}
