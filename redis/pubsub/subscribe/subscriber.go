package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Println(err)
	}
	var wg sync.WaitGroup
	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			subscribe(c, "example")
		}
		defer wg.Done()
	}()
	wg.Wait()
}

func subscribe(c redis.Conn, ch string) {
	psc := redis.PubSubConn{Conn: c}
	psc.Subscribe(ch)
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
		}
	}
}
