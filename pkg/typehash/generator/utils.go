package generator

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"text/template"

	"github.com/denis96z/go-typehash/pkg/typehash"
)

func GenerateCodeForType(v any, tpName string) string {
	b, err := typehash.MD5(v)
	if err != nil {
		panic(err)
	}

	s := "func (v " + tpName + ") TypeHash() string {\n"
	s += fmt.Sprintf("return %q\n", hex.EncodeToString(b))
	s += "}"

	return s
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
