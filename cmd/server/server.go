package main

import (
	"golang-gin-starter/cmd/server/app"
	"os"
)

func main()  {
	command := app.NewCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
