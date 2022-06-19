package base_redis_test

import (
	"fmt"
	"github.com/charlie-bit/base-redis"
	"testing"
)

func TestSGet(t *testing.T) {
	fmt.Println(base_redis.SGet())
}
