package recommend

import (
	"github.com/hicsgo/glib"
	"encoding/json"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 达观智能推荐模块
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type recommend struct {
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 上报数据到达观
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (r *recommend) Upload(request *RecommendRequest) (resp *RecommendResponse, err error) {
	url := "http://datareportapi.datagrand.com/data/YOUR_APP_NAME"
	recommendResult := RecommendResponse{}
	result, err := glib.HttpPostJson(url, request)
	if err == nil {
		err = json.Unmarshal(result, &recommendResult)
		if err == nil {
			return &recommendResult, nil
		}
	}
	return nil, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 从达观拉取数据
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (r *recommend) GetDataFromDatagrand() {

}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 新建推荐实例
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewRecommend() *recommend {
	r := new(recommend)
	return r
}
