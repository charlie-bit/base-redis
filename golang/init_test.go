package golang_test

import (
	"fmt"
	"testing"

	"github.com/charlie-bit/base-redis/golang"
)

func TestStandAloneClient(t *testing.T) {
	fmt.Println(golang.StandAloneClient())
}
