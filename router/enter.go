package router

type RouterGroup struct {
	BlogRouter
	ShopRouter
	VoucherRouter
	UserRouter
	UploadRouter
	FollowRouter
	CommentRouter
}

var RouterGroupApp = new(RouterGroup)
