package main

import (
	"fmt"
	"os"
)

const (
	helpStr = `
Your static site cocktail maker.

usage: go-blender [options]

option:
	init  PATH    initialize default go-blender project in PATH
	build PATH    builds project in currect directory
`
)

// main parses arguments and builds website
func main() {
	logsInit()
	args := os.Args

	if len(args) < 2 {
		InfoLogger.Println("Too few arguments")
		fmt.Println(helpStr)
		return
	}

	switch args[1] {
	case "init":
		if len(args) < 3 {
			fmt.Println(helpStr)
			return
		}
		initPath := args[2]
		blenderInit(initPath)
	case "build":
		if len(args) < 3 {
			fmt.Println(helpStr)
			return
		}
		buildPath := args[2]
		blenderBuild(buildPath)
	default:
		fmt.Println(helpStr)
	}

}
