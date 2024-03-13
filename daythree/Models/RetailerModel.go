package Models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ProductID   string `json:"id" gorm:"type:varchar(50)"`
	ProductName string `json:"product_name" gorm:"type:varchar(100)"`
	Price       uint   `json:"price" gorm:"type:int"`
	Quantity    uint16 `json:"quantity" gorm:"type:int"`
}

type Customer struct {
	gorm.Model
	CustomerID string `json:"customer_id" gorm:"type:varchar(50)"`
	Name       string `json:"customer_name" gorm:"type:varchar(150)"`
}

type Orders struct {
	gorm.Model
}
