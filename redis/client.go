package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var (
	REDISADDR     = "localhost:6379"
	REDISPASSWORD = ""
	REDISDB       = 8

	SESSIONKEY = "UserSession"
)

func NewClient() *redis.Client {
	opt := &redis.Options{
		Addr:     REDISADDR,
		Password: REDISPASSWORD,
		DB:       REDISDB,
	}
	client := redis.NewClient(opt)
	return client
}

func SetSessionKey(username, value string) error {
	client := NewClient()

	if err := client.HSet(SESSIONKEY, username, value).Err(); err != nil {
		return err
	}
	return nil
}

func GetSessionKey(username string) string {
	client := NewClient()
	str, err := client.HGet(SESSIONKEY, username).Result()
	if err != nil {
		log.Fatal(err)
	}
	return str
}

func main() {
	username := "zwhset"
	SetSessionKey(username, "fuck you gg.")
	v := GetSessionKey("zwhset")
	fmt.Printf("[%s] := %s\n", username, v)
}
