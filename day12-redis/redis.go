package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	c, err := redis.Dial("tcp", ":6379")

	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.Do("SET", "message1", "yello humans!")

	world, err := redis.String(c.Do("GET", "message1"))
	if err != nil {
		fmt.Println("key not found yo")
		panic(err)
	}

	fmt.Println(world)
}
