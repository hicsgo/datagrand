package recommend

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * action_type说明
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
const (
	VIEW      = "view"      //点击进入物品详情页。对于资讯，表示进入资讯阅读页面；对于短视频，表示进入播放页；对于直播，表示进入直播间；对于小说，表示进入书籍详情页等等
	READING   = "reading"   //阅读，仅对小说类场景使用，表示进入章节阅读页面
	PLAY      = "play"      //播放，一般对于视频类场景使用，表示对视频进行播放
	REC_CLICK = "rec_click" //点击达观推荐结果，此时最好将此次推荐结果对应的request_id上报到user_action的rec_requestid中
	REC_SHOW  = "rec_show"  //用户浏览过达观推荐结果，此时最好将此次推荐结果对应的request_id上报到user_action的rec_requestid中。主要标示用户浏览了哪些推荐结果
	COLLECT   = "collect"   //收藏
	SUBSCRIBE = "subscribe" //订阅
	COMMENT   = "comment"   //评论
	GIFT      = "gift"      //送礼物
	SHARE     = "share"     //分享
	LIKE      = "like"      //点赞
	DISLIKE   = "dislike"   //点衰
	CART      = "cart"      //加入购物车或加入书架
	BUY       = "buy"       //购买
	SEARCH    = "search"    //搜索
	FOLLOW    = "follow"    //关注
	UNFOLLOW  = "unfollow"  //取消关注
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * table_name说明
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
const (
	ITEM        = "item"        //物品信息表
	USER_ACTION = "user_action" //用户行为数据表
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * cmd说明
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
const (
	ADD         = "add"         //新增或更新一条记录，如果主键对应的记录已经存在，则对该记录做更新操作
	UPDATE      = "update"      //更新一条记录，如果主键对应的记录不存在，则不处理，user_action表不支持update操作
	DELETE      = "delete"      //删除一条记录，如果主键对应的记录不存在，则认为删除成功，user_action表不支持delete操作
	REFRESH_ALL = "refresh_all" //批量下线item数据，仅对item表有效
)
