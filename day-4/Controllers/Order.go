package Controllers

import (
	"bootcamp-prac/day-4/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//OrderProduct ... Order product for the user
func OrderProduct(c *gin.Context){
	var ord Models.Order
	err := c.BindJSON(&ord)
	if err != nil {
		return
	}
	possible := isOrderPossible(ord,c)
	if possible == false{
		c.JSON(http.StatusOK,gin.H{
			"status":"order failed",
		})
		return
	}
	cooldown := isCoolDownOver()
	if cooldown == false{
		c.JSON(http.StatusOK,gin.H{
			"status":"cannot make another order request for next 1 minute",
		})
		return
	}
	err = Models.OrderProduct(&ord)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id": ord.ID,
			"product_id": ord.ProductID,
			"quantity": ord.Quantity,
			"status": "order placed",
		})
	}
}

func GetOrderByID(c *gin.Context){
	id := c.Params.ByName("id")
	var ord Models.Order
	err := Models.GetOrderByID(&ord,id)
	if err !=nil {
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK,ord)
	}
}