package recommend

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 达观请求参数说明
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type RecommendRequest struct {
	AppId        int64         `form:"appid" json:"appid"`                 //您的达观账号对应的id
	TableName    string        `form:"table_name" json:"table_name"`       //要上传数据的表名(item、user_action)
	TableContent []*CmdRequest `form:"table_content" json:"table_content"` //上报的数据内容，JSON格式
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * cmd请求参数说明
 * eg:1
 * {
 * 	"appid": "youappid",
 * 	"table_name": "item",
 * 	"table_content": [{
 * 		"cmd": "add",
 * 		"fields": {
 * 			"itemid": "28394556",
 * 			"cateid": "9_2_1",
 * 			"score": 459,
 * 			"title": "天穿修炼记最新版",
 * 			"item_tags": "修仙"
 * 		}
 * 	}]
 * }
 * eg:2
 * {
 * 	"appid": "",
 * 	"table_name": "user_action",
 * 	"table_content": [{
 * 		"cmd": "add",
 * 		"fields": {
 * 			"timestamp": 1489040079,
 * 			"action_type": "view",
 * 			"userid": "1234",
 * 			"itemid": "abc12"
 * 		}
 * 	}]
 * }
 * eg:3
 * {
 * 	"cmd": "refresh_all",
 * 	"fields": {
 * 		"all_itemid": ["28394556", "23423423"]
 * 	}
 * }
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type CmdRequest struct {
	Cmd    string      `form:"cmd" json:"cmd"`       //要对数据进行的操作(update/delete/add/refresh_all)
	Fields interface{} `form:"fields" json:"fields"` //根据table_name、cmd两个字段确定具体内容
}
