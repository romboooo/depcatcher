package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newInspectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "inspect <path>",
		Short: "извлечение требований из elf-файла",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]

			if _, err := os.Stat(path); err != nil {
				return fmt.Errorf("cannot access %q: %w", path, err)
			}

			inspectElfFile(path)
			return nil
		},
	}
}

func inspectElfFile(path string) {
	fmt.Println(path)
}
