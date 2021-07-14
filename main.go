package main

import (
	"os"

	"Github.com/NaujOyamat/microservice-template/core"
)

func main() {
	args := []string{}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	core.BuildWebHost(args, func() core.IStartup {
		return &Startup{}
	}).Run()
}
