package resolver

import (
	"fmt"
	"path/filepath"
	"strings"
)

func ExpandPaths(binaryPath string, entries []string) ([]string, error) {
	absPath, err := filepath.Abs(binaryPath)

	if err != nil {
		return nil, err
	}

	realPath, err := filepath.EvalSymlinks(absPath)
	if err != nil {
		return nil, err
	}

	origin := filepath.Dir(realPath)

	var paths []string
	for _, entry := range entries {
		for _, path := range strings.Split(entry, ":") {
			if path == "" {
				return nil, fmt.Errorf("empty search path is not supported")
			}

			path = strings.ReplaceAll(path, "${ORIGIN}", origin)
			path = strings.ReplaceAll(path, "$ORIGIN", origin)

			paths = append(paths, filepath.Clean(path))
		}
	}

	return paths, nil
}
