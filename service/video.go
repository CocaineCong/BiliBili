package service

import (
	"BiliBili.com/cache"
	"BiliBili.com/model"
	"BiliBili.com/utils"
	"strconv"
)

func ClicksStoreInDB() {
	utils.Logfile("[info]", " Clicks are stored in the database")
	var vid int          //视频id
	var key string       //redis的key
	var clicks int       //点击量数字
	var strClicks string //字符串格式
	videos := cache.RedisClient.LRange(cache.ClicksVideoList, 0, -1).Val()
	for _, i := range videos {
		vid, _ = strconv.Atoi(i)
		key = cache.VideoClicksKey(vid)
		strClicks, _ = cache.RedisClient.Get(key).Result()
		clicks, _ = strconv.Atoi(strClicks)
		//删除redis数据
		cache.RedisClient.Del(key)
		//写入数据库
		model.DB.Model(&model.Video{}).Where("id = ?", vid).Update("clicks", clicks)
	}
	//删除list
	cache.RedisClient.Del(cache.ClicksVideoList)
	utils.Logfile("[info]", " Click volume storage completed")
}

