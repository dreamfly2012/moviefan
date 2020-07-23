package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	var a = 3
	var b = 10
	res := sum(a, b)
	fmt.Printf("%d 与%d之和:为%d", a, b, res)
	if res != 13 {
		t.Error("error")
	}
}
