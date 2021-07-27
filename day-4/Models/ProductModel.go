package Models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	//ProductId   uint   `gorm:"primaryKey" json:"product_id"`
	ProductName string `json:"product_name" gorm:"column:product_name"`
	Price       int    `json:"price" gorm:"column:price"`
	Quantity    int    `json:"quantity" gorm:"column:quantity"`
}
func (b *Product) TableName() string {
	return "PRODUCTS"
}
//gorm:"primary_key" gorm:"autoIncrement"