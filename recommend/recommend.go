package recommend

import (
	"github.com/hicsgo/glib"
	"encoding/json"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 达观智能推荐模块
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type recommend struct {
	prefixUrl string
	appName   string
	appid     string
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 上报数据到达观
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (r *recommend) Upload(request *RecommendRequest) (resp *RecommendResponse, err error) {
	recommendResult := RecommendResponse{}
	url := r.prefixUrl + r.appName
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
func NewRecommend(prefixUrl, appname, appid string) *recommend {
	r := new(recommend)
	r.prefixUrl = prefixUrl
	r.appName = appname
	r.appid = appid
	return r
}
