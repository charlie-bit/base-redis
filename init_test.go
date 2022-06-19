package base_redis_test

import (
	"fmt"
	"github.com/charlie-bit/base-redis"
	"testing"
)

func TestStandAloneClient(t *testing.T) {
	fmt.Println(base_redis.StandAloneClient())
}

func TestClusterClient(t *testing.T) {
	fmt.Println(base_redis.ClusterClient())
}

func TestSentinelClient(t *testing.T) {
	fmt.Println(base_redis.SentinelClient())
}
