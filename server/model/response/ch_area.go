package response

import "gin-vue-admin/model"

type ChAreaResponse struct {
	Area model.ChArea `json:"area"`
}
