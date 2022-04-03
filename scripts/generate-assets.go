//usr/bin/env go run $0 $@ ; exit
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Working with assets from %s\n", pwd)
	// Convert all .svg into .png from assets directory
	svgFiles, err := filepath.Glob("*.svg")
	if err != nil {
		panic(err)
	}
	for _, svg := range svgFiles {
		fmt.Printf("Processing asset %s ...", svg)
		c := exec.Command("rsvg-convert", "-f", "png", "-o", strings.ReplaceAll(svg, ".svg", ".png"), svg)
		if b, err := c.CombinedOutput(); err != nil {
			panic(err)
		} else {
			fmt.Print(string(b))
		}
		fmt.Println()
	}
}
