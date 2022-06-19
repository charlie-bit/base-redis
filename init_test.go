package golang_test

import (
	"fmt"
	"testing"

	"github.com/charlie-bit/base-redis/golang"
)

func TestStandAloneClient(t *testing.T) {
	fmt.Println(golang.StandAloneClient())
}

func TestClusterClient(t *testing.T) {
	fmt.Println(golang.ClusterClient())
}

func TestSentinelClient(t *testing.T) {
	fmt.Println(golang.SentinelClient())
}
