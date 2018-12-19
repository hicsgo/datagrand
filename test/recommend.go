package test

import (
	"testing"
	"datagrand/recommend"
	"fmt"
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 测试上传数据到达观
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func TestGetSensorsData(t *testing.T) {
	r := recommend.NewRecommend()
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
	result, err := r.Upload(request)
	fmt.Println(&result, err)
}

func A(t *testing.T) {
	t.Name()
}
