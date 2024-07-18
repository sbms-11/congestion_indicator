package router

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	//user用画面 エンドポイント
	r.GET("/", controller.Get_getShopStatusList)
	r.GET("/home", controller.Get_getShopInfo)
	r.GET("/myPage", controller.Post_makeRes)
	r.GET("/shopStatus", controller.Post_makeRes)
	r.POST("/MakeNewRes", controller.Post_deleteRes)

	//shop用画面 エンドポイント
	r.GET("/shop_home", controller.Get_getResStatus)
	r.GET("/shop_makeNewRes", controller.Post_makeResForShop)
	r.GET("/shop", controller.Post_deleteResForShop)
	r.GET("/shop", controller.Post_makeEnterdStatus)
}
