package resolver

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func FindLibrary(name string, dirs []string) (string, error) {

	for _, dir := range dirs {
		candidate := filepath.Join(dir, name)
		info, err := os.Stat(candidate)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return "", fmt.Errorf("err with fining a lib: %w", err)
		}
		if !info.IsDir() {
			return candidate, nil
		}
	}

	return "", nil
}
