[![GitHub go.mod Go version of a Go module](https://img.shields.io/badge/go-1.23.3-blue)](https://go.dev/dl/)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/badge/work_with-redis-red)](https://go.dev/dl/)
[![Go Reference](https://pkg.go.dev/badge/github.com/rabbitmq/amqp091-go.svg)](https://pkg.go.dev/github.com/gopy-art/zrediss)


# zrediss
simple tool for interacting with redis databases and redis databases with SSL

## Quick Start
first you have to download the package with this command :
```sh
go get github.com/gopy-art/zrediss
```
then you can use this package in your code, here is an example of using this package: 
```go
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
```

## Start With Module
first you have to download the project or clone it from the github. 

then go to the root directory of the project and run this command : 
```sh
go build
```

now you cna use this module in your server !

get all the keys : 
```sh
./zrediss -u redis://localhost:6379 -action all
```

get one key with value : 
```sh
./zrediss -u redis://localhost:6379 -action get -k key1
```

set one key with value : 
```sh
./zrediss -u redis://localhost:6379 -action set -k key1 -s value1
```

delete all keys from redis : 
```sh
./zrediss -u redis://localhost:6379 -action flush
```

delete just one key : 
```sh
./zrediss -u redis://localhost:6379 -action delete -k key1
```

## falgs
```txt
  -action string
        set the action of the module. (all, get, set, delete, flush)
  -k string
        set rediss key to get the value
  -l string
        set app logger type , stdout or file (default "stdout")
  -s string
        set value for the specific key
  -u string
        set rediss url
  -v    zrediss version
  -val string
        set rediss value for set the value
```
