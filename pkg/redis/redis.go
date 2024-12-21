package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var Client *redis.Client

// ConnectRedis initializes the Redis client
func ConnectRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "redis-15027.c1.ap-southeast-1-1.ec2.redns.redis-cloud.com:15027", // Địa chỉ Redis server
		Username: "default",                                                         // Tên đăng nhập Redis (nếu có)
		Password: "5osKjJZZejs0QtFNzErWyC61AtUSXA0a",                                // Mật khẩu Redis (nếu có)
		DB:       0,                                                                 // Sử dụng database Redis mặc định
	})

	ctx := context.Background()

	// Kiểm tra kết nối Redis
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis!")
}
