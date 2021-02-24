package main

import (
	"golang-gin-starter/cmd/server/app"
	"math/rand"
	"os"
	"time"
)

func main()  {
	// 每次使用不同的种子来渠道不同的随机数
	rand.Seed(time.Now().UnixNano())

	command := app.NewCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
