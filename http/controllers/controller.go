package controllers

import (
	"assignment2/model"
	"assignment2/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type controller struct {
	repo repository.OrdersRepo
}

func NewOrderByController(repo repository.OrdersRepo) *controller {
	return &controller{repo}
}

func (cnt *controller) CreateOrderBy(c *gin.Context) {
	var param model.Order
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := cnt.repo.Create(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": result})
}

func (cnt *controller) GetAllOrderBy(c *gin.Context) {
	result, err := cnt.repo.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": result})
}

func (cnt *controller) UpdateOrderBy(c *gin.Context) {
	var param model.Order
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := cnt.repo.Update(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": result})
}

func (cnt *controller) DeleteOrderBy(c *gin.Context) {
	id := c.Param("id")

	idint, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = cnt.repo.Delete(int64(idint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
