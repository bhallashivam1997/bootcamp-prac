package Controllers

import (
	"bootcamp-prac/day-4/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func isOrderPossible(ord Models.Order, c *gin.Context) bool{
	id := ord.ProductID
	var prod Models.Product
	err := Models.GetProductByID(&prod,strconv.Itoa(id))
	if err !=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}
	if prod.Quantity < ord.Quantity{
		return false
	}
	prod.Quantity -= ord.Quantity
	err = Models.UpdateProduct(&prod,strconv.Itoa(id))
	if err !=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}
	return true
}