package typehash

import (
	"crypto/md5"
	"hash"
	"io"
	"reflect"
	"sort"
)

func MD5(v any) ([]byte, error) {
	return calculateHash(md5.New(), v)
}

func calculateHash(ctx hash.Hash, v any) ([]byte, error) {
	updateCtxFromAny(ctx, reflect.TypeOf(v))
	return ctx.Sum(nil), nil
}

func updateCtxFromAny(ctx hash.Hash, tp reflect.Type) {
	switch tp.Kind() {
	case reflect.Struct:
		updateCtxFromStruct(ctx, tp)
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Pointer:
		updateCtxFromAny(ctx, tp.Elem())
	default:
		updateCtxFromSimple(ctx, tp)
	}
}

func updateCtxFromStruct(ctx hash.Hash, tp reflect.Type) {
	n := tp.NumField()

	fields := make(structFieldSlice, 0, n)
	for i := 0; i < tp.NumField(); i++ {
		f := tp.Field(i)
		fields = append(
			fields, structField{
				Name: f.Name,
				Type: f.Type,
			},
		)
	}

	sort.Sort(fields)
	for i := 0; i < n; i++ {
		mustWriteString(ctx, fields[i].Name)
		updateCtxFromAny(ctx, fields[i].Type)
	}
}

type structField struct {
	Name string
	Type reflect.Type
}

type structFieldSlice []structField

func (v structFieldSlice) Len() int {
	return len(v)
}

func (v structFieldSlice) Less(i, j int) bool {
	return v[i].Name < v[j].Name
}

func (v structFieldSlice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func updateCtxFromSimple(ctx hash.Hash, tp reflect.Type) {
	mustWriteString(ctx, tp.Name())
}

func mustWriteString(w io.Writer, s string) {
	b := []byte(s)

	n, err := w.Write(b)
	if err != nil {
		panic(err)
	}
	if n != len(b) {
		panic("not all bytes written")
	}
}
