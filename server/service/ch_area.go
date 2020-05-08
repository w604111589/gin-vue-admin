package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

//GetArea ...
func GetArea(areaID uint) (area model.ChArea, err error) {
	err = global.GVA_DB.Table("ch_area").Where("f_area_id = ?", areaID).First(&area).Error
	return
}

//GetAreaList ...
func GetAreaList() (arealist []model.ChArea, err error) {
	err = global.GVA_DB.Table("ch_area").Find(&arealist).Error
	return
}

//GetAreaListByParentID 通过parentID获取子地区列表
func GetAreaListByParentID(parentID int) (areaList []model.ChArea, err error) {
	err = global.GVA_DB.Table("ch_area").Where("f_parent_id = ?", parentID).Find(&areaList).Error
	return
}
