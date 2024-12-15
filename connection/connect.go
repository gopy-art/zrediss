package connection

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
)

/*
main struct for interact with redis database
*/
type RedisConnection struct {
	RedisAddress    string
	RedisConnection *redis.Client
}

/*
This method will make a connection to the redis database and save the pointer of the connection in the main struct
*/
func (r *RedisConnection) InitConnection() error {
	addr, err := redis.ParseURL(r.RedisAddress)
	if err != nil {
		return fmt.Errorf("the url of redis server is invalid")
	}

	if strings.Contains(r.RedisAddress, "rediss://") {
		r.RedisConnection = redis.NewClient(&redis.Options{
			Addr:      addr.Addr,
			Password:  addr.Password,
			DB:        addr.DB,
			TLSConfig: &tls.Config{InsecureSkipVerify: true},
		})
	} else {
		r.RedisConnection = redis.NewClient(&redis.Options{
			Addr:     addr.Addr,
			Password: addr.Password,
			DB:       addr.DB,
		})
	}

	return nil
}

/*
This method get all the keys from the redis and return them as a []string
*/
func (r *RedisConnection) GetAllKeys() ([]string, error) {
	ctx := context.Background()

	result, err := r.RedisConnection.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}

/*
This method get just one key with value from the data base and it will return the value of that key with an error
*/
func (r *RedisConnection) GetSpecificKey(key string) (string, error) {
	ctx := context.Background()

	val, err := r.RedisConnection.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("key %s does not exist", key)
		} else {
			return "", err
		}
	}

	return val, nil
}

/*
This method flush all the database and delete all the keys from the redis, and it will return 'ok' if every thing run successfully
*/
func (r *RedisConnection) DeleteAllKeys() (string, error) {
	ctx := context.Background()
	err := r.RedisConnection.FlushDB(ctx).Err()
	if err != nil {
		return "", err
	}

	return "ok", nil
}

/*
This method set a key with value in the redis, and it will return 'ok' if every thing run successfully
*/
func (r *RedisConnection) SetKeyWithValue(key, value string) (string, error) {
	ctx := context.Background()

	_, err := r.RedisConnection.Set(ctx, key, value, 0).Result()
	if err != nil {
		return "", err
	}

	return "ok", nil
}

/*
This method delete just one key from redis database, and it will return 'ok' if every thing run successfully
*/
func (r *RedisConnection) DeleteSpecificKey(key string) (string, error) {
	ctx := context.Background()

	err := r.RedisConnection.Del(ctx, key).Err()
	if err != nil {
		return "", fmt.Errorf("error deleting key %s: %v", key, err)
	}

	return "ok", nil
}
