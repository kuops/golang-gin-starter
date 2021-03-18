package app

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"golang-gin-starter/internal/pkg/server"
)


func NewCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "server",
		Short: "Run gin-starter server",
		Long: "The gin-starter is a golang gin framework template",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello gin starter!")
			server.Run(context.Background())
		},
	}
	return command
}
