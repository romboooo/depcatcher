package cli

import (
	"fmt"
	"os"

	"github.com/romboooo/depcatcher/internal/elfscan"
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

	libs := res.Libraries
	runPaths := res.RunPath
	rPaths := res.RPath
	if err != nil {
		return err
	}
	fmt.Println("libs:")
	for i := range libs {
		fmt.Printf("  %s\n", libs[i])
	}

	fmt.Printf("architecture: %s\n", res.Machine)
	fmt.Printf("class: %s\n", res.Class)
	fmt.Printf("interpreter:%s\n", res.Interpreter)
	fmt.Println("runpaths:")
	for i := range runPaths {
		fmt.Printf("  %s\n", runPaths[i])
	}
	fmt.Println("rPaths:")
	for i := range rPaths {
		fmt.Printf("  %s\n", rPaths[i])
	}

	return nil
}
