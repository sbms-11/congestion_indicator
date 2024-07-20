package main

import (
	"github.com/sbms-11/congestion_indicator/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.ApiRouter(r)
	r.Run()
}
