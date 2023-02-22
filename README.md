# go-typehash

Work In Progress!

## Example

go.mod
```
module example

go 1.20
```

main.go
```go
package main

import (
	"fmt"

	"example/types"
)

func main() {
	v1 := types.T1{}
	fmt.Println(v1.TypeHash())

	v2 := types.T2{}
	fmt.Println(v2.TypeHash())

	fmt.Println(v1.TypeHash() == v2.TypeHash())
}
```

types/data.go
```go
package types

//go:generate typehash-gen

//typehash:md5
type T1 struct {
	A1 string
	A2 string
}

//typehash:md5
type T2 struct {
	A2 string
	A1 string
}

type T3 struct {
	B int
}
```
