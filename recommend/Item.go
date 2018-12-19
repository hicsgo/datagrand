package recommend

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 达观物品信息结构
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type Item struct {
	ItemId         string `form:"itemid" json:"itemid"`                     //item的唯一id
	CateId         string `form:"cateid" json:"cateid"`                     //有则上报，有助于提升推荐效果	item所属分类，多级分类用“_”进行分隔。如：“1_2”,表示一级分类为1，二级分类为2。如果同时属于多个分类，用英文分号“;”分隔,如：“1_2;3_1”
	Score          int64  `form:"score" json:"score"`                       //score	int	有则上报，有助于提升推荐效果	item的质量分，0~1000的整数。值越大，被推荐出来的可能性越大
	Title          string `form:"title" json:"title"`                       //有则上报，有助于提升推荐效果	item的标题
	Content        string `form:"content" json:"content"`                   //正文，比如帖子或新闻的正文
	Price          int64  `form:"price" json:"price"`                       //有则上报，有助于提升推荐效果	item的价格，单位为分
	ItemTags       string `form:"item_tags" json:"item_tags"`               //有则上报，有助于提升推荐效果	item的标签信息，多个标签以英文分号“;”分割。如“变形金钢5;擎天柱;大黄蜂”，item_tags一般是一组主题关键词。
	ItemModifyTime int64  `form:"item_modify_time" json:"item_modify_time"` //有则上报，有助于提升推荐效果，同时可以控制推荐结果的时效性	item最新修改时间，unix时间戳，精确到秒。上报该字段可以提高推荐的时效性。
	DgStartTime    int64  `form:"dg_start_time" json:"dg_start_time"`       //item开始推荐的时间，unix时间戳，精确到秒。有值的话，没到此时间则不能被推荐出来。无值的话，开始推荐的时间不受限制
	DgEndTime      int64  `form:"dg_end_time" json:"dg_end_time"`           //item结束推荐的时间，unix时间戳，精确到秒。有值的话，超过此时间则不能被推荐出来。无值的话，结束推荐的时间不受限制
	Publisher      string `form:"publisher" json:"publisher"`               //有则上报，有助于提升推荐效果	item的上传者信息，如可以对应视频或者文章的发布者、小说的作者等
	Province       string `form:"province" json:"province"`                 //如果有地域相关需求，有则上报	item所属的地域信息，这里指省级地域，如江苏省、广东省等
	City           string `form:"city" json:"city"`                         //如果有地域相关需求，有则上报	item所属的地域信息，这里指市级地域，如苏州市、杭州市等/
	Extend         string `form:"extend" json:"extend"`                     //自定义扩展字段
}
