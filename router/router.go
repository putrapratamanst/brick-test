package router

import (
	controller "brick-test/controller/product"
	"brick-test/service/product"

	"github.com/gin-gonic/gin"
)

func Route(v1 *gin.RouterGroup, prd *product.Service) {
	handler := controller.ControllerHandler(prd)
	product := v1.Group("/product")
	{
		product.GET("/scrap", handler.Scrap)
	}
}