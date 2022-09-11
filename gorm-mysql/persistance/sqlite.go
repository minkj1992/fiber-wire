package persistance

import (
	"github.com/k0kubun/pp/v3"
	"github.com/minkj1992/go-playground/gorm-mysql/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqlite() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&domain.Product{})

	// Create
	db.Create(&domain.Product{Code: "D42", Price: 100})

	// Read
	var product domain.Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(domain.Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	pp.Println(product)
	// Delete - delete product
	db.Delete(&product, 1)
}
