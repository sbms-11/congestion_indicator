package main

import (
	"backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.ApiRouter(r)
	r.Run()
}
