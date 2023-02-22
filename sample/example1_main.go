//go:build example1

package main

import (
	"fmt"

	"github.com/denis96z/go-typehash/sample"
)

func main() {
	v1 := sample.T1{}
	fmt.Println(v1.TypeHash())

	v2 := sample.T2{}
	fmt.Println(v2.TypeHash())

	fmt.Println(v1.TypeHash() == v2.TypeHash())
}
