package golang

import (
	"fmt"
	"github.com/go-redis/redis"
)

/**
	redis 连接方式存在三种
	1. 单机连接
	2. 集群连接
	3. 哨兵模式连接
**/

// StandAloneClient 单机模式
// 参数详解
/* type Options struct {
	// 网络类型 tcp 或者 unix.
	// 默认是 tcp.
	Network string
	// redis地址，格式 host:port
	Addr string

	// 新建一个redis连接的时候，会回调这个函数
	OnConnect func(*Conn) error

	// redis密码，redis server没有设置可以为空。
	Password string

	// redis数据库，序号从0开始，默认是0，可以不用设置
	DB int

	// redis操作失败最大重试次数，默认不重试。
	MaxRetries int

	// 最小重试时间间隔.
	// 默认是 8ms ; -1 表示关闭.
	MinRetryBackoff time.Duration

	// 最大重试时间间隔
	// 默认是 512ms; -1 表示关闭.
	MaxRetryBackoff time.Duration

	// redis连接超时时间.
	// 默认是 5 秒.
	DialTimeout time.Duration

	// socket读取超时时间
	// 默认 3 秒.
	ReadTimeout time.Duration

	// socket写超时时间
	WriteTimeout time.Duration

	// redis连接池的最大连接数.
	// 默认连接池大小等于 cpu个数 * 10
	PoolSize int

	// redis连接池最小空闲连接数.
	MinIdleConns int
	// redis连接最大的存活时间，默认不会关闭过时的连接.
	MaxConnAge time.Duration

	// 当你从redis连接池获取一个连接之后，连接池最多等待这个拿出去的连接多长时间。
	// 默认是等待 ReadTimeout + 1 秒.
	PoolTimeout time.Duration
	// redis连接池多久会关闭一个空闲连接.
	// 默认是 5 分钟. -1 则表示关闭这个配置项
	IdleTimeout time.Duration
	// 多长时间检测一下，空闲连接
	// 默认是 1 分钟. -1 表示关闭空闲连接检测
	IdleCheckFrequency time.Duration

	// 只读设置，如果设置为true， redis只能查询缓存不能更新。
	readOnly bool
} */
func StandAloneClient() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	ok, err := rdb.Ping().Result()
	if err != nil {
		fmt.Printf("init redis failed, err: %s \n", err.Error())
		return err
	}
	fmt.Printf("redis ping result : %s \n", ok)
	return nil
}
