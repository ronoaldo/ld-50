package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	fmt.Println("Building wasm ... ")
	cmd := exec.Command("go", "build", "-o", "ld-50.wasm", "../../")
	cmd.Env = append(os.Environ(), "GOOS=js")
	cmd.Env = append(cmd.Env, "GOARCH=wasm")
	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", string(b))
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Println("Copying wasm_exec.js from GOROOT")
	b, err = exec.Command("go", "env", "GOROOT").CombinedOutput()
	if err != nil {
		fmt.Printf("%s", string(b))
		fmt.Printf("Error: %v", err)
		return
	}
	goroot := strings.TrimSpace(string(b))
	in, err := ioutil.ReadFile(path.Join(goroot, "misc", "wasm", "wasm_exec.js"))
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	ioutil.WriteFile("wasm_exec.js", in, 0666)
}
