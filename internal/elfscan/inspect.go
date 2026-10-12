package elfscan

import (
	"debug/elf"
	"fmt"
	"io"
	"strings"
)

type Result struct {
	Machine     elf.Machine
	Class       elf.Class
	Libraries   []string
	Interpreter string
	RunPath     []string
	RPath       []string
}

func Inspect(path string) (Result, error) {
	file, err := elf.Open(path)
	if err != nil {
		return Result{}, fmt.Errorf("open ELF %q: %w", path, err)
	}
	defer file.Close()

	interpreter, err := findInterpreter(file)

	if err != nil {
		return Result{}, err
	}
	runPaths, err := file.DynString(elf.DT_RUNPATH)

	if err != nil {
		return Result{}, fmt.Errorf("read RUNPATH of %q: %w", path, err)
	}

	rPaths, err := file.DynString(elf.DT_RPATH)
	if err != nil {
		return Result{}, fmt.Errorf("read RPATH of %q: %w", path, err)
	}
	libs, err := file.ImportedLibraries()
	if err != nil {
		return Result{}, fmt.Errorf("read dependencies of %q: %w", path, err)
	}

	return Result{
		Machine:     file.Machine,
		Class:       file.Class,
		Libraries:   libs,
		Interpreter: interpreter,
		RunPath:     runPaths,
		RPath:       rPaths,
	}, nil
}

func findInterpreter(file *elf.File) (string, error) {
	progs := file.Progs
	var res string
	for _, prog := range progs {
		if prog.Type != elf.PT_INTERP {
			continue
		}

		data, err := io.ReadAll(prog.Open())

		if err != nil {
			return "", err
		}

		res = strings.TrimSuffix(string(data), "\x00")
		break
	}

	return res, nil

}
