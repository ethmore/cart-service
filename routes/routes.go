package routes

import (
	"cart/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/test", controllers.Test())
	g.POST("/addProductToCart", controllers.AddToCart())
	g.POST("/getCartProducts", controllers.GetCartProducts())
	g.POST("/removeProductFromCart", controllers.RemoveProductFromCart())
	g.POST("/changeProductQty", controllers.ChangeProductQty())
	g.POST("/toPurchase", controllers.ToPurchase())
	g.POST("/getTotalPrice", controllers.GetTotalPrice())
}
