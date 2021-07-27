//Controllers/Product.go
package Controllers
import (
	"bootcamp-prac/day-4/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
//GetProducts ... Get all products
func GetProducts(c *gin.Context) {
	var prod []Models.Product
	err := Models.GetAllProducts(&prod)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, prod)
	}
}
//CreateProduct ... Create Product
func CreateProduct(c *gin.Context) {
	var prod Models.Product
	err := c.BindJSON(&prod)
	if err != nil {
		return
	}
	err = Models.CreateProduct(&prod)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id": prod.ID,
			"product_name": prod.ProductName,
			"price": prod.Price,
			"quantity": prod.Quantity,
			"message": "product successfully added",
		})
	}
}
//GetProductByID ... Get the product by id
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var prod Models.Product
	err := Models.GetProductByID(&prod, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK,prod)
	}
}
//UpdateProduct ... Update the user information
func UpdateProduct(c *gin.Context) {
	var user Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	err = c.BindJSON(&user)
	if err != nil {
		return
	}
	err = Models.UpdateProduct(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//DeleteProduct ... Delete the user
func DeleteProduct(c *gin.Context) {
	var prod Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&prod, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
