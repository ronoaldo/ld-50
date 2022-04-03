//usr/bin/env go run $0 $@ ; exit
package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Convert all .svg into .png from assets directory
	svgFiles, err := filepath.Glob("assets/*.svg")
	if err != nil {
		panic(err)
	}
	for _, svg := range svgFiles {
		c := exec.Command("rsvg-convert", "-f", "png", "-o", strings.ReplaceAll(svg, ".svg", ".png"), svg)
		if b, err := c.CombinedOutput(); err != nil {
			panic(err)
		} else {
			fmt.Print(string(b))
		}
	}
}
