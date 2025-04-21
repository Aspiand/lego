package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouters(router *gin.Engine) {
	registerProductRoutes(router)
	registerBrandRoutes(router)
}
