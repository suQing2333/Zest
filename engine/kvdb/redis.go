package kvdb

import (
	"fmt"
	"github.com/gomodule/redigo/redis" //引入redis
	"time"
	"zest/engine/common"
	"zest/engine/conf"
	"zest/engine/zslog"
)

type redisConf struct {
	ip         string
	port       int
	password   string
	idle_time  time.Duration
	max_idle   int
	max_active int
	uri        string
}

const (
	REDIS_IP         = "127.0.0.1"
	REDIS_PORT       = 6379
	REDIS_PW         = ""
	REDIS_IDLE_TIME  = 100
	REDIS_MAX_IDLE   = 8
	REDIS_MAX_ACTIVE = 0
)

var pool *redis.Pool
var redisConfInfo *redisConf

func init() {
	loadRedisConfig()
	ConnectRedis()
}

func loadRedisConfig() {
	redisConfInfo = new(redisConf)
	ip := common.ThreeUnary(conf.IsSet("redis.ip"), conf.GetString("redis.ip"), REDIS_IP)
	port := common.ThreeUnary(conf.IsSet("redis.port"), conf.GetInt("redis.port"), REDIS_PORT)
	password := common.ThreeUnary(conf.IsSet("redis.password"), conf.GetString("redis.password"), REDIS_PW)
	idleTime := common.ThreeUnary(conf.IsSet("redis.idle_time"), conf.GetInt("redis.idle_time"), REDIS_IDLE_TIME)
	maxIdle := common.ThreeUnary(conf.IsSet("redis.max_idle"), conf.GetInt("redis.max_idle"), REDIS_MAX_IDLE)
	maxActive := common.ThreeUnary(conf.IsSet("redis.max_active"), conf.GetInt("redis.max_active"), REDIS_MAX_ACTIVE)
	uri := fmt.Sprintf("%v:%v", ip, port)

	redisConfInfo.ip = common.String(ip)
	redisConfInfo.port = common.Int(port)
	redisConfInfo.password = common.String(password)
	redisConfInfo.idle_time = time.Duration(common.Int(idleTime))
	redisConfInfo.max_idle = common.Int(maxIdle)
	redisConfInfo.max_active = common.Int(maxActive)
	redisConfInfo.uri = common.String(uri)
}

func ConnectRedis() {
	setPassword := redis.DialPassword(redisConfInfo.password)

	pool = &redis.Pool{
		MaxIdle:     redisConfInfo.max_idle,
		MaxActive:   redisConfInfo.max_active,
		IdleTimeout: redisConfInfo.idle_time,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redisConfInfo.uri, setPassword)
			if err != nil {
				zslog.LogError("connect redis error %v", err)
				return nil, err
			}
			return conn, err
		},
	}
	zslog.LogDebug("connect redis pool successs")
}

func CloseRedis() {
	err := pool.Close()
	if err != nil {
		zslog.LogError("%v", err)
	}
	zslog.LogDebug("close redis pool successs")
}

func Get(args ...interface{}) string {
	conn := pool.Get()
	cmd := "get"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Set(args ...interface{}) string {
	conn := pool.Get()
	cmd := "set"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func GetSet(args ...interface{}) string {
	conn := pool.Get()
	cmd := "getset"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Append(args ...interface{}) string {
	conn := pool.Get()
	cmd := "append"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Incrby(args ...interface{}) string {
	conn := pool.Get()
	cmd := "incrby"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Mset(args ...interface{}) string {
	conn := pool.Get()
	cmd := "mset"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Msetnx(args ...interface{}) string {
	conn := pool.Get()
	cmd := "msetnx"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Mget(args ...interface{}) string {
	conn := pool.Get()
	cmd := "mget"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Hset(args ...interface{}) string {
	conn := pool.Get()
	cmd := "hset"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Hsetnx(args ...interface{}) string {
	conn := pool.Get()
	cmd := "hsetnx"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Hget(args ...interface{}) {
	conn := pool.Get()
	cmd := "hget"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Hexists(args ...interface{}) {
	conn := pool.Get()
	cmd := "hexists"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Hdel(args ...interface{}) string {
	conn := pool.Get()
	cmd := "hdel"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Hlen(args ...interface{}) string {
	conn := pool.Get()
	cmd := "hlen"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Hincrby(args ...interface{}) string {
	conn := pool.Get()
	cmd := "hincrby"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Hmset(args ...interface{}) string {
	conn := pool.Get()
	cmd := "hmset"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Hmget(args ...interface{}) string {
	conn := pool.Get()
	cmd := "hmget"
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%v err = %v", cmd, err)
		return ""
	}
	return fmt.Sprintf("%v", ret)
}

func Hkeys(args ...interface{}) {}

func Hvals(args ...interface{}) {}

func Hgetall(args ...interface{}) {}

func Sadd(args ...interface{}) {}

func Sismember(args ...interface{}) {}

func Spop(args ...interface{}) {}

func Srandmember(args ...interface{}) {}

func Srem(args ...interface{}) {}

func Smove(args ...interface{}) {}

func Scard(args ...interface{}) {}

func Smembers(args ...interface{}) {}

func Sinter(args ...interface{}) {}

func Sinterstore(args ...interface{}) {}

func Sunion(args ...interface{}) {}

func Sunionstore(args ...interface{}) {}

func Sdiff(args ...interface{}) {}

func Sdiffstore(args ...interface{}) {}

func Zadd(args ...interface{}) {}

func Zscore(args ...interface{}) {}

func Zincrby(args ...interface{}) {}

func Zcard(args ...interface{}) {}

func Zcount(args ...interface{}) {}

func Zrange(args ...interface{}) {}

func Zrevrange(args ...interface{}) {}

func Zrangebyscore(args ...interface{}) {}

func Zrevrangebyscore(args ...interface{}) {}

func Zrank(args ...interface{}) {}

func Zrevrank(args ...interface{}) {}

func Zrem(args ...interface{}) {}

func Zremrangebyrank(args ...interface{}) {}

func Zremrangebyscore(args ...interface{}) {}

func Zunionstore(args ...interface{}) {}

func Zinterstore(args ...interface{}) {}

func Redis(cmd string, args ...interface{}) string {
	conn := pool.Get()
	defer conn.Close()

	ret, err := conn.Do(cmd, args...)
	if err != nil {
		fmt.Printf("%s err= %s", cmd, err)
		return ""
	}
	return fmt.Sprintf("%s", ret)
}
