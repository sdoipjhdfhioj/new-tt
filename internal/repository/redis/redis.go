package redis

import (
	"github.com/go-redis/redis/v7"

	"sync"
)

var red *Redis

var once sync.Once

func InitRedis(addr string) *Redis {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr: addr,
		})
		r := client.Ping()
		if r.Err() != nil {
			panic(r.Err())
		}
		red = &Redis{client: client}
	})

	return red
}

type Redis struct {
	client *redis.Client
}

func (p *Redis) Set(key string, value interface{}) error {
	r := p.client.Set(key, value, 0)
	if r.Err() != nil {
		return r.Err()
	}
	return nil
}

func (p *Redis) Get(key string) (string, error) {
	r := p.client.Get(key)
	if r.Err() != nil {
		return "", r.Err()
	}
	return r.Val(), nil
}

func (p *Redis) Remove(key string) error {
	r := p.client.Del(key)
	if r.Err() != nil {
		return r.Err()
	}
	return nil
}
