package vo

type UserVo struct {
	ID       uint64 `json:"id,string"`
	NickName string `json:"nickName"`
	Icon     string `json:"icon"`
}
