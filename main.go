package main

import (
	"os"

	"github.com/awalvie/go-blender/cli"
	"github.com/awalvie/go-blender/logging"
)

// main parses arguments and builds website
func main() {
	logging.Init()
	args := os.Args

	if len(args) < 2 {
		cli.Help()
		return
	}

	switch args[1] {
	case "init":
		if len(args) < 3 {
			logging.ErrorLogger.Fatalln("init: too few arguments")
		}

		initPath := args[2]
		err := cli.Init(initPath)
		if err != nil {
			logging.ErrorLogger.Fatalln("init: ", err)
		}

	case "build":
		if len(args) < 3 {
			logging.ErrorLogger.Fatalln("build: too few arguments")
		}
		buildPath := args[2]
		err := cli.Build(buildPath)
		if err != nil {
			logging.ErrorLogger.Fatalln("build: ", err)
		}

	default:
		cli.Help()
		return
	}

}
