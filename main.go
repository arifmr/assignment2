package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Item struct {
	ItemID      int
	ItemCode    int
	Description string
	Quantity    int
	OrderId     int
}

type Order struct {
	OrderID      int
	CustomerName string `gorm:"type:varchar(100)"`
	OrderedAt    time.Time

	Items Item
}

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	db    *gorm.DB
	errDb error
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=orders_by port=5432 sslmode=disable"
	db, errDb = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDb != nil {
		fmt.Println("Error: ", errDb)
	}
	db.Debug().AutoMigrate(Order{}, Item{})

	r := gin.Default()
	p := r.Group("/orders")
	{
		p.GET("/", Get)
		p.POST("/", Create)
		p.PUT("/", Update)
		p.DELETE("/:id", Delete)
	}
	r.Run()
}

func Get(c *gin.Context) {
	var orders []Order
	result := db.Find(&orders)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func Create(c *gin.Context) {
	var order Order
	c.BindJSON(&order)

	result := db.Create(&order)
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func Update(c *gin.Context) {
	// var order Order

}

func Delete(c *gin.Context) {
	var order Order
	db.Delete(&order)
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully."})
}
