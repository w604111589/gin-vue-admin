package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

//GetArea ...
func GetArea(areaID uint) (area model.ChArea, err error) {
	err = global.GVA_DB.Table("ch_area").Where("f_area_id = ?", areaID).First(&area).Error
	// err = global.GVA_DB.Where("f_area_id = ?", areaID).First(&area).Error
	return
}
