package cli

import (
	"fmt"
	"os"

	"github.com/romboooo/depcatcher/internal/elfscan"
	"github.com/romboooo/depcatcher/internal/resolver"
	"github.com/spf13/cobra"
)

func newInspectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "inspect <path>",
		Short: "извлечение требований из elf-файла",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			cmd.SilenceUsage = true
			if _, err := os.Stat(path); err != nil {
				return fmt.Errorf("cannot access %q: %w", path, err)
			}

			return inspectElfFile(path)
		},
	}
}

func inspectElfFile(path string) error {
	res, err := elfscan.Inspect(path)
	if err != nil {
		return err
	}

	searchEntries := res.RunPath
	if len(searchEntries) == 0 {
		searchEntries = res.RPath
	}

	paths, err := resolver.ExpandPaths(path, searchEntries)
	if err != nil {
		return err
	}
	paths = append(paths, "/usr/lib")
	resolved := make(map[string]string, len(res.Libraries))
	for _, lib := range res.Libraries {
		found, err := resolver.FindLibrary(lib, paths)
		if err != nil {
			return err
		}
		resolved[lib] = found
	}

	writeResult(res, paths, resolved)

	return nil
}

func writeResult(res elfscan.Result, paths []string, resolved map[string]string) {
	fmt.Println("libs:")
	for _, lib := range res.Libraries {
		fmt.Printf("  %s\n", lib)
	}

	fmt.Printf("architecture: %s\n", res.Machine)
	fmt.Printf("class: %s\n", res.Class)
	fmt.Printf("interpreter:%s\n", res.Interpreter)

	fmt.Println("runpaths:")
	for _, rp := range res.RunPath {
		fmt.Printf("  %s\n", rp)
	}

	fmt.Println("rPaths:")
	for _, rp := range res.RPath {
		fmt.Printf("  %s\n", rp)
	}

	fmt.Println("expanded paths:")
	for _, p := range paths {
		fmt.Printf("  %s\n", p)
	}

	fmt.Println("resolved libraries:")
	for _, lib := range res.Libraries {
		found := resolved[lib]
		if found == "" {
			found = "not found"
		}
		fmt.Printf("  %s: %s\n", lib, found)
	}
}
