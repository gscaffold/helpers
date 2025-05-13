package rest

type (
	PageBaseRequest struct {
		Limit   int    `form:"limit" json:"limit"`
		Offset  int    `form:"offset" json:"offset"`
		LastID  int64  `form:"last_id" json:"last_id"`
		Order   string `form:"order" json:"order"`       // default desc.
		OrderBy string `form:"order_by" json:"order_by"` // default id
	}
	PageBaseResponse struct {
		Total   int         `json:"total"`
		Records interface{} `json:"records"`
		LastID  interface{} `json:"last_id"`
	}
)

func (req *PageBaseRequest) LoadDefault() {
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.Order == "" {
		req.Order = "desc"
	}
	if req.OrderBy == "" {
		req.OrderBy = "id"
	}
}
