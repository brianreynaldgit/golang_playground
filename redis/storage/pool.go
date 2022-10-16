package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

//make sure to install redis on docker locally using :
// docker run --name local-redis -p 6379:6379 -d redis
func InitPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   0,
		MaxActive: 1,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6379")

			// Connection error handling
			if err != nil {
				log.Printf("ERROR: fail initializing the redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
}

func InitPool2() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   50,
		MaxActive: 10000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6379")

			// Connection error handling
			if err != nil {
				log.Printf("ERROR: fail initializing the redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
}

func Ping(c redis.Conn) error {
	s, err := redis.String(c.Do("PING"))
	if err != nil {
		return err
	}

	fmt.Printf("PING Response = %s\n", s)

	return nil
}

func Set(c redis.Conn) error {
	_, err := c.Do("SET", "name", "rommel")
	if err != nil {
		return err
	}

	_, err = c.Do("SET", "country", "Philippines")
	if err != nil {
		return err
	}

	return nil
}

func Get(c redis.Conn) error {
	key := "name"
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		return (err)
	}
	fmt.Printf("%s = %s\n", key, s)

	key = "country"
	i, err := redis.String(c.Do("GET", key))
	if err != nil {
		return (err)
	}

	fmt.Printf("%s = %s\n", key, i)

	// Example where key was not set
	key = "Nonexistent Key"
	s, err = redis.String(c.Do("GET", key))
	if err == redis.ErrNil {
		fmt.Printf("%s : Alert! this Key does not exist\n", key)
	} else if err != nil {
		return err
	} else {
		fmt.Printf("%s = %s\n", key, s)
	}

	return nil

}
