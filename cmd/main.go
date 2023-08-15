package main

import (
	"Calculator/src/infrastructure/container"
	"os"
)

func main() {
	container := container.NewContainer()
	container.GetCli().Run()
	os.Exit(0)
}
