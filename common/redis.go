package common

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var RDB *Redis

type IRedis interface {
	Read(key string) (string, error)
	Write(key string, value string) (bool, error)
	Connect(addr string, password string, db string) *redis.Client
}

type Redis struct {
	Addr     string
	Password string
	Db       int
	Client   *redis.Client
}

func (r *Redis) Connect() *redis.Client {
	if r.Client != nil {
		return r.Client
	}
	r.Client = redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password, // no password set
		DB:       r.Db,       // use default DB
	})
	pong, err := r.Client.Ping(ctx).Result()
	log.Print(pong, err)

	RDB = r

	return r.Client
}

func (r *Redis) Read(key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	return val, err
}

func (r *Redis) Write(key string, value string, expired time.Duration) (bool, error) {
	err := r.Client.Set(ctx, key, value, expired).Err()
	return err == nil, err
}
