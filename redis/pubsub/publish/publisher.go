package main

import (
	"time"

	redigo "github.com/brianreynaldgit/golang_playground/redis/storage"
)

func main() {
	redisPool := redigo.InitPool()

	c := redisPool.Get()
	ch := "example"
	c.Do("PUBLISH", ch, "hello "+time.Now().String())
}
