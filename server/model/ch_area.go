package model

type ChArea struct {
	AreaID      uint   `json:"areaId" form:"areaId" gorm:"primary_key;column:f_area_id"`
	ParentID    int    `json:"parentId" form:"parentId" gorm:"column:f_parent_id"`
	AreaName    string `json:"areaName" form:"areaName" gorm:"column:f_area_name"`
	ALLName     string `json:"allName" form:"allName" gorm:"column:f_all_name"`
	FirstLetter string `json:"firstLetter" form:"firstLetter" gorm:"column:f_first_letter"`
	CityJianpin string `json:"cityJianpin" form:"cityJianpin" gorm:"column:f_first_letter"`
	CityQp      string `json:"cityQp" form:"cityQp" gorm:"column:f_city_qp"`
	CityLevel   int    `json:"cityLevel" form:"cityLevel" gorm:"column:f_city_level"`
	CityMsg     string `json:"cityMsg" form:"cityMsg" gorm:"column:f_city_status"`
	CityStatus  string `json:"cityStatus" form:"cityStatus" gorm:"column:f_city_status"`
	ProvinceID  string `json:"provinceId" form:"provinceId" gorm:"column:f_province_id"`
	CityID      string `json:"cityId" form:"cityId" gorm:"column:f_city_id"`
	CountyID    string `json:"countyId" form:"countyId" gorm:"column:f_county_id"`
}
