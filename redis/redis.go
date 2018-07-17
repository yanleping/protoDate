package redis

import (
	//"mobvista.com/adn/zeusextractor/util"
	"sync"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	conn     redis.Conn
	connGet  redis.Conn
	getmutex sync.RWMutex
	host   string
)

// 初始化，存入ip:port 地址 list
func InitRedis(host string) error {
	host = host
	if err := InitSetConn(host); err != nil {
		return err
	}
	if err := InitGetConn(host); err != nil {
		return err
	}
	return nil
}

func InitSetConn(host string) error {
	if conn != nil {
		conn.Close()
	}
	var err error
	fmt.Println(host)
	conn, err = redis.Dial("tcp", host)
	return err
}

func InitGetConn(host string) error {
	if connGet != nil {
		connGet.Close()
	}
	var err error
	fmt.Println(host)
	connGet, err = redis.Dial("tcp", host)
	return err
}

// set  key value 超时时间，失败则返回err
func RedisSet(tablename, key, value string, retrytime int) error {
	_, err := conn.Do("HSET", tablename, key, value)
	if err != nil && retrytime > 0 {
		InitSetConn(host)
		_, err = conn.Do("HSET", tablename, key, value)
		return err
	}
	return err
}

func RedisInfo(retrytime int) (val string, err error) {
	getmutex.Lock()
	defer getmutex.Unlock()
	val, err = redis.String(connGet.Do("INFO"))
	if err != nil && retrytime > 0 {
		InitGetConn(host)
		return "", err
	}
	return val, err

}

// hdel 删除 field，失败则返回err
func RedisHDel(tablename, key string, retrytime int) error {
	_, err := conn.Do("HDEL", tablename, key)
	if err != nil && retrytime > 0 {
		InitRedis(host)
		_, err = conn.Do("HDEL", tablename, key)
		return err
	}
	return err
}

func RedisHKeys(key string, retrytime int) (keys interface{}, err error) {
	keys, err = conn.Do("HKEYS", key)
	if err != nil && retrytime > 0 {
		InitRedis(host)
		keys, err = conn.Do("HKEYS", key)
		return keys, err
	}
	return keys, err
}

// del 删除 key，失败则返回err
func RedisDel(key string, retrytime int) error {
	_, err := conn.Do("DEL", key)
	if err != nil && retrytime > 0 {
		InitRedis(host)
		_, err = conn.Do("DEL", key)
		return err
	}
	return err
}

// get key，返回val，失败则返回err
func RedisGet(tablename, key string, retrytime int) (val string, err error) {
	getmutex.Lock()
	defer getmutex.Unlock()
	val, err = redis.String(connGet.Do("HGET", tablename, key))
	if err != nil && retrytime > 0 {
		InitGetConn(host)
		return "", err
	}
	return val, err
}

// 批量set，如果key value的size不同，取最小的size
func RedisSetBatch(tablename string, key string, value string) error {
	conn.Send("HSET", tablename, key, value)
	return nil
}

// 批量del，如果key value的size不同，取最小的size
func RedisDelBatch(tablename string, key string) error {
	conn.Send("HDEL", tablename, key)
	return nil
}

// 操作
func RedisSend(commandName string, args ...interface{}) error {
	conn.Send(commandName, args...)
	return nil
}

func RedisFlush(keynum int) error {
	conn.Flush()
	var err error
	for i := 0; i < keynum; i++ {
		_, curerr := conn.Receive()
		if curerr != nil {
			err = curerr
		}
	}
	return err
}
