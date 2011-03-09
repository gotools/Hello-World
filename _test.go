package reverse_test

import (
	"fmt"
	"github.com/gotools/reverse"
	"testing"
)

func Test0(t *testing.T) {
	in := "hello world! this is test0."
	fmt.Println(in)
	fmt.Println("reverse:", reverse.Reverse(in))
}

