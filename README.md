# Redis client for Go

![build workflow](https://github.com/go-redis/redis/actions/workflows/build.yml/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-redis/redis/v8)](https://pkg.go.dev/github.com/go-redis/redis/v8?tab=doc)
[![Documentation](https://img.shields.io/badge/redis-documentation-informational)](https://redis.uptrace.dev/)

> refer document, sort basic golang demo for redis.
> Give it a star as well!

## Resources

- [Redis Document](https://www.redis.com.cn/tutorial.html)
- [Reference](http://www.lsdcloud.com/go/middleware/go-redis.html)
- [Chat](554486596@qq.com)

## Features

- Redis basic commands
- Automatic connection pooling with
- [Pub/Sub](https://redis.uptrace.dev/guide/go-redis-pubsub.html).
- [Pipelines and transactions](https://redis.uptrace.dev/guide/go-redis-pipelines.html).
- [Scripting](https://redis.uptrace.dev/guide/lua-scripting.html).
- [Cron Job]().
- [Basic Definition]().

## Installation

go-redis supports 2 last Go versions and requires a Go version with
[modules](https://github.com/golang/go/wiki/Modules) support. So make sure to initialize a Go module:

```shell
go mod init github.com/my/repo
```

If you are using **Redis 6**, install go-redis/**v8**:

```shell
go get github.com/go-redis/redis/v8
```

If you are using **Redis 7**, install go-redis/**v9**:

```shell
go get github.com/go-redis/redis/v9
```

## Quickstart

```go
import (
"context"
"fmt"

"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ExampleClient() {
rdb := redis.NewClient(&redis.Options{
Addr:     "localhost:6379",
Password: "", // no password set
DB:       0, // use default DB
})

err := rdb.Set(ctx, "key", "value", 0).Err()
if err != nil {
panic(err)
}

val, err := rdb.Get(ctx, "key").Result()
if err != nil {
panic(err)
}
fmt.Println("key", val)

val2, err := rdb.Get(ctx, "key2").Result()
if err == redis.Nil {
fmt.Println("key2 does not exist")
} else if err != nil {
panic(err)
} else {
fmt.Println("key2", val2)
}
// Output: key value
// key2 does not exist
}
```

## Look and feel

Some corner cases:

```go
// SET key value EX 10 NX
set, err := rdb.SetNX(ctx, "key", "value", 10*time.Second).Result()

// SET key value keepttl NX
set, err := rdb.SetNX(ctx, "key", "value", redis.KeepTTL).Result()

// SORT list LIMIT 0 2 ASC
vals, err := rdb.Sort(ctx, "list", &redis.Sort{Offset: 0, Count: 2, Order: "ASC"}).Result()

// ZRANGEBYSCORE zset -inf +inf WITHSCORES LIMIT 0 2
vals, err := rdb.ZRangeByScoreWithScores(ctx, "zset", &redis.ZRangeBy{
Min: "-inf",
Max: "+inf",
Offset: 0,
Count: 2,
}).Result()

// ZINTERSTORE out 2 zset1 zset2 WEIGHTS 2 3 AGGREGATE SUM
vals, err := rdb.ZInterStore(ctx, "out", &redis.ZStore{
Keys: []string{"zset1", "zset2"},
Weights: []int64{2, 3}
}).Result()

// EVAL "return {KEYS[1],ARGV[1]}" 1 "key" "hello"
vals, err := rdb.Eval(ctx, "return {KEYS[1],ARGV[1]}", []string{"key"}, "hello").Result()

// custom command
res, err := rdb.Do(ctx, "set", "key", "value").Result()
```

## Run the test

go-redis will start a redis-server and run the test cases.

The paths of basic-redis file and test files are defined in `XXXX_test.go`:

```go
go test -v golang/set_test.go
```

For local testing, you can change the variables to refer to your local files, or test for specific method

```go
go test -v golang/set_test.go -test.run TestSSet
```
