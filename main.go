package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
)

var (
	URL   string
	All   bool
	Key   string
	Flush bool
	Set   string
	Value string
)

func main() {
	// set flags
	flag.StringVar(&URL, "u", "", "set rediss url")
	flag.StringVar(&Key, "k", "", "set rediss key to get the value")
	flag.StringVar(&Set, "s", "", "set rediss key to set the value")
	flag.StringVar(&Value, "v", "", "set rediss value for set the value")
	flag.BoolVar(&All, "a", false, "get all of the redis keys")
	flag.BoolVar(&Flush, "f", false, "delete all of the redis keys")
	flag.Parse()

	// connect to the redis
	var rdb *redis.Client
	addr, err := redis.ParseURL(URL)
	if err != nil {
		log.Fatal("the url of redis server is invalid")
	}

	if strings.Contains(URL, "rediss://") {
		rdb = redis.NewClient(&redis.Options{
			Addr:      addr.Addr,
			Password:  addr.Password,
			DB:        addr.DB,
			TLSConfig: &tls.Config{InsecureSkipVerify: true},
		})
	} else {
		rdb = redis.NewClient(&redis.Options{
			Addr:     addr.Addr,
			Password: addr.Password,
			DB:       addr.DB,
		})
	}

	// get all of the keys
	if All {
		ctx := context.Background()

		// Get all keys starting with *
		result, err := rdb.Keys(ctx, "*").Result()
		if err != nil {
			panic(err)
		}

		if len(result) == 0 {
			fmt.Println("redis in empty!")
		}

		for n, key := range result {
			fmt.Printf("%v ) %v \n", n, key)
		}
		os.Exit(2)
	}

	if Key != "" {
		ctx := context.Background()

		val, err := rdb.Get(ctx, Key).Result()
		if err != nil {
			if err == redis.Nil {
				fmt.Printf("key %s does not exist\n", Key)
			} else {
				panic(err)
			}
		} else {
			fmt.Printf("'%v' : %v\n", Key, val)
		}
	}

	if Flush {
		ctx := context.Background()
		err := rdb.FlushDB(ctx).Err()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("ok")
		}
	}

	if Set != "" && Value != "" {
		ctx := context.Background()

		_, err := rdb.Set(ctx, Set, Value, 0).Result()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ok")
		}
	}
}
