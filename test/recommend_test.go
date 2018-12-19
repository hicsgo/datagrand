package test

import (
	"testing"
	"datagrand/recommend"
	"fmt"
	"encoding/json"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 测试上传数据到达观
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func TestGetSensorsData(t *testing.T) {
	r := recommend.NewRecommend("http://datareportapi.datagrand.com/data/", "your_app_name", "your_app_id")

	request := &recommend.RecommendRequest{
		AppId:     999,
		TableName: recommend.ITEM,
		TableContent: []*recommend.CmdRequest{
			&recommend.CmdRequest{
				Cmd: recommend.ADD,
				Fields: &recommend.Item{
					ItemId:   "1111",
					CateId:   "9_8_1",
					Score:    100,
					Title:    "测试测试",
					ItemTags: "修仙",
				},
			},
			&recommend.CmdRequest{
				Cmd: recommend.ADD,
				Fields: &recommend.Item{
					ItemId:   "1111",
					CateId:   "9_8_1",
					Score:    100,
					Title:    "测试测试1",
					ItemTags: "修仙1",
				},
			},
		},
	}
	requestJson, _ := json.Marshal(request)
	fmt.Println("request is ", string(requestJson))

	result, err := r.Upload(request)

	responseJson, _ := json.Marshal(result)
	fmt.Println("result is ", string(responseJson))

	fmt.Println(result, err)
}
