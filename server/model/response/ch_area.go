package response

import "gin-vue-admin/model"

type ChAreaResponse struct {
	Area model.ChArea `json:"detail"`
}

type ChAreaListResponse struct {
	List []model.ChArea `json:"list"`
}
