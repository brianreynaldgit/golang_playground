package main

import (
	redigo "github.com/brianreynaldgit/golang_playground/redis/storage"
)

//make sure to install redis on docker locally using :
// docker run --name local-redis -p 6379:6379 -d redis
func main() {
	//tried setting first pool
	redisPool := redigo.InitPool()

	conn := redisPool.Get()
	redigo.Ping(conn)
	redigo.Set(conn)
	redigo.Get(conn)

	//tried opening second pool and get value. see if it shares the first pool data (since the connection is basically the same)
	redisPool2 := redigo.InitPool2()
	conn2 := redisPool2.Get()
	redigo.Get(conn2)
}
