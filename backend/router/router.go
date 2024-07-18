package router

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	//user用画面 エンドポイント
	r.GET("/user_home", controller.Get_getShopStatusList)
	r.GET("/user_ShopDetails", controller.Get_getShopDetails)
	r.POST("/user_makeRes", controller.Post_makeRes)
	r.POST("/user_deleteRes", controller.Post_deleteRes)
	r.POST("/user_mypage", controller.Post_makeRes)

	//shop用画面 エンドポイント
	r.GET("/shop_home", controller.Get_getResStatus)
	r.POST("/shop_makeNewRes", controller.Post_makeResForShop)
	r.POST("/shop_deleteRes", controller.Post_deleteResForShop)
	r.POST("/shop_makeEntrydStatus", controller.Post_makeEntrydStatus)
}
