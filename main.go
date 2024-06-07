package main

import (
	"log"

	"github.com/awalvie/go-blender/cmd"
)

func init() {
	// Configure logger flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
func main() {
	cmd.Execute()
}
