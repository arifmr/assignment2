package model

import "gorm.io/gorm"

type (
	Item struct {
		gorm.Model
		ItemCode    string `gorm:"column:item_code; type:varchar(50)"`
		Description string `gorm:"column:description; type:varchar(255)"`
		Quantity    int    `gorm:"column:quantity"`
		OrderID     int64  `gorm:"coulmn:order_id"`
	}

	Order struct {
		gorm.Model
		CustomerName string `gorm:"column:customer_name; type:varchar(50)"`
		Items        []Item `gorm:"foreignKey:order_id; references:ID"`
	}
)
