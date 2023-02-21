package sample

//go:generate typehash-gen

//typehash:md5
type T1 struct {
	A1 string
	A2 string `json:"a2"`
}

type T2 struct {
	B int
}
