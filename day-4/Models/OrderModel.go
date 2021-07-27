package Models

import "github.com/jinzhu/gorm"

type Order struct{
	gorm.Model
	//Id        int    `json:"order_id" gorm:"column:order_id"`
	ProductID int    `json:"product_id" gorm:"column:product_id"`
	Quantity  int    `json:"quantity" gorm:"column:quantity"`
	Status    string `json:"status" gorm:"column:status"`
}

func (b *Order) TableName() string {
	return "ORDERS"
}
//gorm:"primary_key" gorm:"autoIncrement"