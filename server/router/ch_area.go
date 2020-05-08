package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitChAreaRouter(Router *gin.RouterGroup) {
	// ApiRouter := Router.Group("area").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	ApiRouter := Router.Group("common")
	{
		// ApiRouter.POST("area", v1.CreateArea)   // 创建地区
		ApiRouter.PUT("area", v1.UpdateArea) // 更新地区
		//ApiRouter.DELETE("area", v1.DeleteArea)                       // 删除地区
		ApiRouter.GET("area", v1.GetArea)                             // 获取单一地区
		ApiRouter.GET("areaList", v1.GetAreaList)                     // 获取地区列表
		ApiRouter.GET("areaListByParentID", v1.GetAreaListByParentID) // 获取子地区列表
	}
}
