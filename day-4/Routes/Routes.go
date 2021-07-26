//Routes/Routes.go
package Routes
import (
	"bootcamp-prac/day-4/Controllers"
	"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/product-api")
	{
		grp1.GET("products", Controllers.GetProducts)
		grp1.POST("product", Controllers.CreateProduct)
		grp1.GET("product/:id", Controllers.GetProductByID)
		grp1.PATCH("product/:id", Controllers.UpdateProduct)
		grp1.DELETE("product/:id", Controllers.DeleteProduct)
	}
	grp2 :=r.Group("/order")
	{
		grp2.POST("",Controllers.OrderProduct)
		grp2.GET(":id",Controllers.GetOrderByID)
	}
	return r
}