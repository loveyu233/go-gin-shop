package utils

type EnterUtils struct {
	JwtUtils        JwtUtils        //JWT创建和验证
	PhoneCodeUtils  PhoneCodeUtils  //手机号码判断和验证码
	StructMapUtils  StructMapUtils  //结构体转mao
	RedisCacheUtils RedisCacheUtils //redis缓存
	RedisIDWorker   RedisIDWorker   //redis全局唯一ID生成
	TimeUtils       TimeUtils       //时间转化
	RedisBlogLiked  RedisBlogLiked  //blogLiked操作
	RedisGeoUtil    RedisGeoUtil    //	redis经纬度操作
	RedisBitmapUtil RedisBitmapUtil //签到
}

var EnterUtilsApp = new(EnterUtils)
