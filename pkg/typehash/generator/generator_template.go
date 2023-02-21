package generator

const (
	GenFileTemplate = `
{{$info := .}}

//go:build typehash
package main

// Code generated by typehash-gen. DO NOT EDIT.

import (
	"log"

	typehash "github.com/denis96z/go-typehash/pkg/typehash/generator"

	{{$info.Pkg.Name}} {{printf "%q" $info.Pkg.Path}}
)

var (
	{{range $idx, $tp := $info.Types}}
		v{{$idx}} = {{$info.Pkg.Name}}.{{$tp.Name}}{}
	{{end}}
)

const (
	GeneratedFilePath = {{printf "%q" $info.GeneratedFilePath}}
)

func main() {
	gData := typehash.GenData{
		PkgName: {{printf "%q" $info.Pkg.Name}},

		Types: make([]typehash.GenDataTypeInfo, 0),
	}
	{{range $idx, $tp := $info.Types}}
		{
			var gdtpInfo typehash.GenDataTypeInfo
			gdtpInfo.FuncCode = typehash.GenerateCodeForType(v{{$idx}}, "{{$tp.Name}}")

			gData.Types = append(gData.Types, gdtpInfo)
		}
	{{end}}
	if err := typehash.WriteTypeHashFile(GeneratedFilePath, gData); err != nil {
		log.Fatal(err)
	}
}
`
)
