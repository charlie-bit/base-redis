package base_redis_test

import (
	"fmt"
	"github.com/charlie-bit/base-redis"
	"testing"
)

func TestSSet(t *testing.T) {
	fmt.Println(base_redis.SSet())
}

func TestHashSet(t *testing.T) {
	fmt.Println(base_redis.HashSet())
}

func TestMapSet(t *testing.T) {
	fmt.Println(base_redis.MapSet())
}

func TestListSet(t *testing.T) {
	fmt.Println(base_redis.ListSet())
}
