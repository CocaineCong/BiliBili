package cache

import (
	"BiliBili.com/pkg/utils"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"),
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		utils.Logfile("[Error]", " redis error "+err.Error())
	}
	RedisClient = client
}

