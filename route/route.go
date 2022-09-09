package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jufianto/microusers/app/users"
)

func SetupRoute(route *gin.Engine) {
	route.POST("/regis", users.HandlerRegisterUser)
}
