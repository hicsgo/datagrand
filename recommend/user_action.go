package recommend

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 达观用户行为结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type UserAction struct {
	UserId       string `form:"userid" json:"userid"`               //有则上报，否则对推荐结果有负面影响    用户登陆后的id
	IMei         string `form:"imei" json:"imei"`                   //有则上报，否则对推荐结果有负面影响    用户的手机IMEI号
	Cid          string `form:"cid" json:"cid"`                     //有则上报，否则对推荐结果有负面影响    用户的cookieid，（userid、imei、cid不能全部为空，否则会对推荐效果有负面影响）
	ItemId       string `form:"itemid" json:"itemid"`               //除了搜索行为之外，此字段不能为空    对应item表的itemid（如果是follow、unfollow行为，该值对应的是关注、取消关注的对象的id）
	ActionType   string `form:"action_type" json:"action_type"`     //用户对item的具体行为，必须是有意义的字符串，上报用户对达观推荐结果的点击请填“rec_click”，用于统计推荐点击率及推荐效果优化
	ActionNum    int64  `form:"action_num" json:"action_num"`       //行为数量。对于购买行为，可以是购买数量
	ActionDetail string `form:"action_detail" json:"action_detail"` //该行为的一些描述信息，该字段可以自行定义，比如登录的话，可以是qq登录或者微信登录，发送短信可以是短信发送成功或者失败等等。在达观数据后台系统会对这些数据进行汇总，以英文分号进行分隔
	SceneType    string `form:"scene_type" json:"scene_type"`       //场景类型，用于标识不同的场景，常用的包括pc_home（PC首页个性推荐）, android_detail（安卓详情页个性化推荐）等。不在上述中的可以自行添加。
	TimeStamp    string `form:"timestamp" json:"timestamp"`         //行为发生的时间，unix时间戳，精确到秒，没有传则置为收到请求的时间
	RecRequestId string `form:"rec_requestid" json:"rec_requestid"` //如果此条行为是由达观推荐带来的，则此字段对应调达观推荐接口返回itemid所带的request_id
	Type         string `form:"type" json:"type"`                   //action_type是search或dislike或follow或unfollow的话
	Value        string `form:"value" json:"value"`                 //action_type是search或是dislike的话，有则需要上报该字段    该字段配合着type字段进行上报，type目前支持cate、publisher、item。如果type是publisher，则value对应的是publisher id，多个的话以英文分号分隔，如”100; 101″；如果type是item，则value对应的是itemid。如果type是cate，则value对应的是类目id，多个类目的话以英文分号分隔，如1_1; 1_2, 另外，action_type为search，表示在1_1, 1_2类目下有搜索行为，action_type为dislike，表示不喜欢类目1_1, 1_2
	KeyWord      string `form:"keyword" json:"keyword"`             //action_type是search或是dislike的话，有则需要上报该字段    该字段配合着type字段进行上报，type目前支持cate、publisher、item。如果type是publisher，则value对应的是publisher id，多个的话以英文分号分隔，如”100; 101″；如果type是item，则value对应的是itemid。如果type是cate，则value对应的是类目id，多个类目的话以英文分号分隔，如1_1; 1_2, 另外，action_type为search，表示在1_1, 1_2类目下有搜索行为，action_type为dislike，表示不喜欢类目1_1, 1_2
	Extend       string `form:"extend" json:"extend"`               //自定义扩展字段
}
