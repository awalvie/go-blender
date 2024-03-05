package main

import (
	"log"
	"os"

	"github.com/awalvie/go-blender/cli"
)

// main parses arguments and builds website
func main() {
	args := os.Args

	if len(args) < 2 {
		cli.Help()
		return
	}

	switch args[1] {
	case "init":
		if len(args) < 3 {
			log.Fatalln("init: too few arguments")
		}

		initPath := args[2]
		err := cli.Init(initPath)
		if err != nil {
			log.Fatalln("init: ", err)
		}

	case "build":
		if len(args) < 3 {
			log.Fatalln("build: too few arguments")
		}
		buildPath := args[2]
		err := cli.Build(buildPath)
		if err != nil {
			log.Fatalln("build: ", err)
		}

	default:
		cli.Help()
		return
	}

}
