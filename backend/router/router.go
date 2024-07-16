package router

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	r.GET("/", controller.SdampleFanc)
	r.GET("/home", controller.GetHomeInfo)
	r.GET("/myPage", controller.GetUserInfo)
	r.GET("/shopStatus", controller.GetShopStatus)
	r.POST("/MakeNewReservation", controller.MakeNewReservation)
}
