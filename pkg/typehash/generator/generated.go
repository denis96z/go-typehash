package generator

import (
	"fmt"
)

func MakeGeneratedFilePath(srcPath string) string {
	bname := srcPath[:len(srcPath)-len(".go")]
	return bname + "_typehash.go"
}

func GoRunGeneratorFile(gfpath string) error {
	if err := executeCommand("go", "run", gfpath); err != nil {
		return err
	}
	return nil
}

type GenData struct {
	PkgName string

	Types []GenDataTypeInfo
}

type GenDataTypeInfo struct {
	FuncCode string
}

func WriteTypeHashFile(fpath string, gData GenData) error {
	if err := createGoFileFromTemplate(fpath, "typehash", FileTemplate, gData); err != nil {
		return fmt.Errorf("failed to write typehash file [path = %q]: %w", fpath, err)
	}
	return nil
}
