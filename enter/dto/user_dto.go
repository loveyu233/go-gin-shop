package dto

import (
	"go-gin-shop/enter/tb"
	"go-gin-shop/enter/vo"
)

type LoginUserDto struct {
	ID       uint64 `json:"id,string" ` // 主键
	Phone    string `json:"phone" `     // 手机号码
	NickName string `json:"nickName" `  // 昵称，默认是用户id
	Icon     string `json:"icon" `      // 人物头像
	Password string `json:"password"`   // 密码，加密存储
	Code     string `json:"code"`       //手机验证码
}

func UserDtoToUserModel(loginuser LoginUserDto) tb.TbUser {
	return tb.TbUser{
		Phone:    loginuser.Phone,
		Password: loginuser.Password,
		NickName: loginuser.NickName,
	}
}
func UserModelToUserVo(userModel tb.TbUser) vo.UserVo {
	return vo.UserVo{
		ID:       userModel.ID,
		NickName: userModel.NickName,
		Icon:     userModel.Icon,
	}
}
