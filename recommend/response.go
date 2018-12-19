package recommend

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 达观输出参数说明
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type RecommendResponse struct {
	Status    string            `form:"status" json:"status"`         //返回状态
	Errors    []*RecommendError `form:"errors" json:"errors"`         //返回错误code和message
	RequestId int64             `form:"request_id" json:"request_id"` //本次请求Id(排错)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 达观错误参数说明
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type RecommendError struct {
	Code    string `form:"code" json:"form"`       //code
	Message string `form:"message" json:"message"` //返回的错误描述
}
