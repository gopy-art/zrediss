package bin

import (
	"fmt"
	"os"

	"github.com/gopy-art/zrediss/connection"
	"github.com/gopy-art/zrediss/console"
	logger "github.com/gopy-art/zrediss/log"
)

/*
This function is for run the module manually with the flag and dependencies.
it has all actions and structs in it and by running the module this function will call.
*/
func RunModuleWithFlag() {
	if console.URL == "" || console.Action == "" {
		logger.ErrorLogger.Println("redis url or redis action is empty!")
	}

	redisHandler := connection.RedisConnection{
		RedisAddress: console.URL,
	}
	err := redisHandler.InitConnection()
	if err != nil {
		logger.ErrorLogger.Fatalf("error in connect to redis, error = %v", err)
	}

	switch console.Action {
	case "all":
		results, err := redisHandler.GetAllKeys()
		if err != nil {
			logger.ErrorLogger.Fatalf("error in get all keys, error = %v", err)
		}

		if len(results) == 0 {
			logger.WarningLogger.Println("redis database is empty")
			os.Exit(0)
		}

		for n, key := range results {
			fmt.Printf("%v ) %v \n", n, key)
		}
	case "get":
		if console.Key == "" {
			logger.ErrorLogger.Fatal("you should set your key!")
		}

		result, err := redisHandler.GetSpecificKey(console.Key)
		if err != nil {
			logger.ErrorLogger.Fatalf("error in get the key, error = %v", err)
		}

		fmt.Printf("'%v' : %v\n", console.Key, result)
	case "set":
		if console.Key == "" || console.Set == "" {
			logger.ErrorLogger.Fatal("you should set the key and value!")
		}

		result, err := redisHandler.SetKeyWithValue(console.Key, console.Set)
		if err != nil {
			logger.ErrorLogger.Fatalf("error in set key and value, error = %v", err)
		}

		fmt.Println(result)
	case "flush":
		result, err := redisHandler.DeleteAllKeys()
		if err != nil {
			logger.ErrorLogger.Fatalf("error in delete all keys, error = %v", err)
		}

		fmt.Println(result)
	case "delete":
		if console.Key == "" {
			logger.ErrorLogger.Fatal("you should set the key!")
		}

		result, err := redisHandler.DeleteSpecificKey(console.Key)
		if err != nil {
			logger.ErrorLogger.Fatalf("error in delete the key, error = %v", err)
		}

		fmt.Println(result)
	default:
		logger.ErrorLogger.Fatal("invalid action!")
	}
}
