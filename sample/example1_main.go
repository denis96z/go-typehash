//go:build example1

package main

import (
	"fmt"

	"github.com/denis96z/go-typehash/sample"
)

func main() {
	v1 := sample.T1{}
	fmt.Println(v1.TypeHash())
}
