package cli

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "depcatcher",
		Short: "утилита для диагностики зависимостей",
		Long:  "cli-утилита для адаптивной диагностики зависимостей единиц развертывания",
	}

	cmd.AddCommand(newInspectCmd())

	return cmd
}
