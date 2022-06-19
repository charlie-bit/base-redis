package golang_test

import (
	"fmt"
	"github.com/charlie-bit/base-redis/golang"
	"testing"
)

func TestSGet(t *testing.T) {
	fmt.Println(golang.SGet())
}
