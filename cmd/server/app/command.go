package app

import (
	"fmt"
	"github.com/spf13/cobra"
)


func NewCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "server",
		Short: "Run gin-starter server",
		Long: "The gin-starter is a golang gin framework template",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello gin starter!")
		},
	}
	return command
}
