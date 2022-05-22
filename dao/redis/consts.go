package redis

// redis 点赞模块key常量    谁给哪个视频点了什么赞
const (
	Prefix                 = "dousheng:"       // 项目key前缀
	KeyVideoFavoriteZSetPF = "video:favorite:" // zset;用户点赞类型
	KeyUserSetPF           = "user:"           // set;记录每个用户点赞的视频
)

// GetRedisKey 给key拼接前缀
func GetRedisKey(key string) string {
	return Prefix + key
}
