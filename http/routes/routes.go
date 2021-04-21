package routes

import (
	"assignment2/http/controllers"
	"assignment2/repository"

	"github.com/gin-gonic/gin"
)

func NewRouter(repo repository.OrdersRepo) *gin.Engine {
	router := gin.Default()

	cont := controllers.NewOrderByController(repo)

	orders := router.Group("/orders")
	{
		orders.POST("/", cont.CreateOrderBy)
		orders.GET("/", cont.GetAllOrderBy)
		orders.PUT("/", cont.UpdateOrderBy)
		orders.DELETE("/:id", cont.DeleteOrderBy)
	}

	return router
}
