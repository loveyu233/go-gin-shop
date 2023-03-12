package service

type EnterServices struct {
	PaginateService Paginate
	BlogService     Blog
	ShopService     Shop
	UserService     User
	VouCherService  VouCher
	FollowService   Follow
	CommentService  Comment
}

var EnterServicesApp = new(EnterServices)
