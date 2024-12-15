package example

import (
	"fmt"
	"log"
	"os"

	"github.com/gopy-art/zrediss/connection"
)

/*
This function connect to the redis database and get all the keys for more actions
*/
func ConnectAndGetAllKeys() {
	redisHandler := connection.RedisConnection{
		RedisAddress: "redis://localhost:6379",
	}
	err := redisHandler.InitConnection()
	if err != nil {
		log.Fatalf("error in connect to redis, error = %v", err)
	}

	results, err := redisHandler.GetAllKeys()
	if err != nil {
		log.Fatalf("error in get all keys, error = %v", err)
	}

	if len(results) == 0 {
		log.Println("redis database is empty")
		os.Exit(0)
	}

	for n, key := range results {
		fmt.Printf("%v ) %v \n", n, key)
	}
}
