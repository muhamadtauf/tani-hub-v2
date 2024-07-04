package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"tani-hub-v2/constant"
	"tani-hub-v2/database"
	"tani-hub-v2/repository"
	"tani-hub-v2/structs"
)

func InsertOrder(c *gin.Context) {
	var order structs.Order

	err := c.ShouldBindJSON(&order)
	if err != nil {
		panic(err)
	}

	Uuid := uuid.New()
	order.Code = Uuid.String()

	var Total float64
	for index, orderDetail := range order.OrderDetail {

		var product structs.Product
		product.Id = orderDetail.ProductId
		repository.GetProductById(database.DbConnection, product)
		err, products := repository.GetProductById(database.DbConnection, product)
		if err != nil {
			panic(err)
		}

		order.OrderDetail[index].Price = products[0].Price
		order.OrderDetail[index].Total = order.OrderDetail[index].Price * float64(orderDetail.Quantity)
		order.OrderDetail[index].OrderCode = Uuid.String()
		Total += order.OrderDetail[index].Total
	}

	order.Total = Total
	order.Status = constant.ACCEPTED

	err = repository.InsertOrder(database.DbConnection, order)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Order",
	})
}

func GetAllOrder(c *gin.Context) {
	var result gin.H

	orders, err := repository.GetAllOrder(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": orders,
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetOrderByCode(c *gin.Context) {
	var order structs.Order
	uuid := c.Param("uuid")

	order.Code = uuid

	var result gin.H

	orders, err := repository.GetOrderByCode(database.DbConnection, order)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": orders,
		}
	}
	c.JSON(http.StatusOK, result)
}
