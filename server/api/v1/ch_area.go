package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
)

//UpdateArea ...
func UpdateArea(c *gin.Context) {

}

func DeleteArea(c *gin.Context) {

}

//GetArea ...
func GetArea(c *gin.Context) {
	var area model.ChArea
	c.ShouldBindQuery(&area)
	fmt.Println(area.AreaID)
	area, err := service.GetArea(area.AreaID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
		return
	}
	response.OkWithData(resp.ChAreaResponse{Area: area}, c)
}

func GetAreaList(c *gin.Context) {
	areaList, err := service.GetAreaList()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取结果失败 %v", err), c)
		return
	}
	response.OkWithData(resp.ChAreaListResponse{List: areaList}, c)
}

func GetAreaListByParentID(c *gin.Context) {
	var area model.ChArea
	c.ShouldBindQuery(&area)
	fmt.Println(area.ParentID)
	areaList, err := service.GetAreaListByParentID(area.ParentID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取结果失败 %v", err), c)
		return
	}
	response.OkWithData(resp.ChAreaListResponse{List: areaList}, c)
}
