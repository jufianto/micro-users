package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jufianto/microusers/route"
)

func main() {
	port := "8989"

	router := gin.New()

	route.SetupRoute(router)
	router.Run(fmt.Sprintf(":%s", port))
}
