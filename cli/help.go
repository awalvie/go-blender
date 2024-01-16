package cli

import (
	"fmt"
	"os"
)

const (
	helpStr = `
Your static site cocktail maker.

Usage: go-blender [options]

option:
	init  PATH    initialize default go-blender project in PATH
	build PATH    builds project in currect directory`
)

// Help prints the usage message to stdout
func Help() {
	fmt.Fprintln(os.Stdout, helpStr)
}
