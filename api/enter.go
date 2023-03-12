package api

type ApiGroup struct {
	BlogApi    BlogApi
	ShopApi    ShopApi
	VoucherApi VoucherApi
	UserApi    UserApi
	UploadApi  UploadApi
	FollowApi  FollowApi
	CommentApi CommentApi
	CaptchaApi CaptchaApi
}

var ApiGroupApp = new(ApiGroup)
