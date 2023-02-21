package generator

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"hash"
	"io"
	"os"
	"os/exec"
	"reflect"
	"text/template"
)

func GenerateCodeForType(v any, tpName string) string {
	ctx := md5.New()
	updateCtxFromAny(ctx, v)

	s := "func (v " + tpName + ") TypeHash() string {\n"
	s += fmt.Sprintf("return %q\n", hex.EncodeToString(ctx.Sum(nil)))
	s += "}"

	return s
}

func updateCtxFromAny(ctx hash.Hash, v any) {
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Struct:
		updateCtxFromStruct(ctx, val)
	default:
		updateCtxFromSimple(ctx, val)
	}
}

func updateCtxFromStruct(ctx hash.Hash, val reflect.Value) {
	for i := 0; i < val.NumField(); i++ {
		updateCtxFromAny(ctx, val.Field(i))
	}
}

func updateCtxFromSimple(ctx hash.Hash, val reflect.Value) {
	mustWriteString(ctx, val.Type().Name())
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

func createGoFileFromTemplate(fpath string, name, ttxt string, data interface{}) error {
	tmpl, err := template.New(name).Parse(ttxt)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err = (tmpl).Execute(&buf, data); err != nil {
		return fmt.Errorf("template %q execute error: %w", name, err)
	}

	fset := token.NewFileSet()

	astf, err := parser.ParseFile(fset, fpath, buf.Bytes(), parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse file content: %w", err)
	}

	f, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf(
			"failed to create file: %w", err,
		)
	}
	defer func() {
		_ = f.Close()
	}()
	if err = format.Node(f, fset, astf); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func executeCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	{
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			`failed to run command: %s`, err,
		)
	}
	return nil
}
