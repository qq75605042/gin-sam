package datasource

import (
	"fmt"
	"gin-sam/conf"
	"log"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var CacheInstance *RedisConn
var cacheLock sync.Mutex

type RedisConn struct {
	pool      *redis.Pool
	showDebug bool
}

func (rds *RedisConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := rds.pool.Get()
	defer conn.Close()
	t1 := time.Now().UnixNano()
	reply, err = conn.Do(commandName, args...)
	if err != nil {
		e := conn.Err()
		if e != nil {
			log.Fatal("redis.do err:", e)
		}
	}
	t2 := time.Now().UnixNano()
	if rds.showDebug {
		log.Printf("[redis] [info] [%dus] cmd=%s,err=%s,args=%v,reply=%s\n", (t2-t1)/1000, commandName, err, args, reply)
	}
	return reply, err
}

func (rds *RedisConn) ShowDebug(b bool) {
	rds.showDebug = b
}

func InstanceCache() *RedisConn {
	if CacheInstance != nil {
		return CacheInstance
	}
	cacheLock.Lock()
	defer cacheLock.Unlock()
	if CacheInstance != nil {
		return CacheInstance
	}
	return NewCache()
}

func NewCache() *RedisConn {
	pool := redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			dial, e := redis.Dial("tcp", fmt.Sprintf("%s:%d", conf.RedisCache.Host, conf.RedisCache.Port))
			if e != nil {
				log.Fatal("redis dial err:", e)
				return nil, e
			}
			_, e = dial.Do("select", 1)
			if e != nil {
				log.Fatal("redis select db err:", e)
				return nil, e
			}
			return dial, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         1000,
		MaxActive:       1000,
		IdleTimeout:     0,
		Wait:            false,
		MaxConnLifetime: 0,
	}
	instance := &RedisConn{
		pool:      &pool,
		showDebug: false,
	}
	CacheInstance = instance
	instance.ShowDebug(true)
	return instance
}
