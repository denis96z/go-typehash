package sample

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
