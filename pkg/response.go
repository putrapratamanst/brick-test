package pkg

import (
	"brick-test/presenter"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, res *presenter.Response) {
	if res.Code == 0 && res.Message == "" {
		res.Code = 204
		res.Message = "No Content"
	}

	ctx.JSON(res.Code, res)
}
