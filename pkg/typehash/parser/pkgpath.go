package parser

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func getPkgPath(fpath string) (string, error) {
	dir := filepath.Dir(fpath)
	for {
		infoArr, err := os.ReadDir(dir)
		if err != nil {
			return "", fmt.Errorf(
				"failed to read directory [path = %q]: %w", dir, err,
			)
		}
		for _, info := range infoArr {
			if info.Name() == "go.mod" {
				mfpath := path.Join(dir, info.Name())

				module, err := getModuleFromGoModFile(mfpath)
				if err != nil {
					return "", err
				}

				module += strings.Join(
					filepath.SplitList(
						filepath.Dir(strings.TrimPrefix(fpath, dir)),
					),
					"/",
				)

				return module, nil
			}
		}
		dir = filepath.Dir(dir)
	}
}

var (
	modRxp = regexp.MustCompile(`^(module\s+)(\S+)$`)
)

func getModuleFromGoModFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf(
			"failed to read go.mod file [path = %q]: %w", path, err,
		)
	}

	lines := strings.Split(string(b), "\n")
	if len(lines) == 0 {
		return "", fmt.Errorf(
			"invalid go.mod file [path = %q]: empty", path,
		)
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)

		m := modRxp.FindStringSubmatch(line)
		if len(m) == 0 {
			continue
		}

		return m[2], nil
	}

	return "", fmt.Errorf(
		"invalid go.mod file [path = %q]: no module name", path,
	)
}
