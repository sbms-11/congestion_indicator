package router

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	//user用画面 エンドポイント
	r.GET("/", controller.Get_getShopStatusList)
	r.GET("/home", controller.Get_getShopInfo)
	r.GET("/myPage", controller.Post_makeReservation)
	r.GET("/shopStatus", controller.Post_makeReservation)
	r.POST("/MakeNewReservation", controller.Post_deleteReservation)

	//shop用画面 エンドポイント
	r.GET("/shop", controller.Get_getReservationStatus)
	r.GET("/shop", controller.Post_makeReservationForShop)
	r.GET("/shop", controller.Post_deleteReservationForShop)
	r.GET("/shop", controller.Post_makeEnterdStatus)
}
