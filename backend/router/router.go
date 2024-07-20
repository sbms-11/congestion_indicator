package router

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	//user用画面 エンドポイント
	userGroup := r.Group("/user")
	{
		userGroup.GET("/user_home", controller.Get_getShopStatusList)
		userGroup.GET("/user_ShopDetails", controller.Get_getShopDetails)
		userGroup.POST("/user_makeRes", controller.Post_makeRes)
		userGroup.POST("/user_deleteRes", controller.Post_deleteRes)
		userGroup.POST("/user_mypage", controller.Post_makeRes)
	}

	//shop用画面 エンドポイント
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/shop_home", controller.Get_getResStatus)
		shopGroup.POST("/shop_makeNewRes", controller.Post_makeResForShop)
		shopGroup.POST("/shop_deleteRes", controller.Post_deleteResForShop)
		shopGroup.POST("/shop_makeEntrydStatus", controller.Post_makeEntrydStatus)
	}
}
