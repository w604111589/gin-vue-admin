package model

type ChArea struct {
	AreaID      uint   `json:"areaId" form:"areaId" gorm:"primary_key"`
	ParentID    int    `json:"parentId" form:"parentId"`
	AreaName    string `json:"areaName" form:"areaName"`
	ALLName     string `json:"allName" form:"allName"`
	FirstLetter string `json:"firstLetter" form:"firstLetter"`
	CityJianpin string `json:"cityJianpin" form:"cityJianpin"`
	CityQp      string `json:"cityQp" form:"cityQp"`
	CityLevel   int    `json:"cityLevel" form:"cityLevel"`
	CityMsg     string `json:"cityMsg" form:"cityMsg"`
	CityStatus  string `json:"cityStatus" form:"cityStatus"`
	ProvinceID  string `json:"provinceId" form:"provinceId"`
	CityID      string `json:"cityId" form:"cityId"`
	CountyID    string `json:"countyId" form:"countyId"`
}
