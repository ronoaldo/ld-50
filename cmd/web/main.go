package main

import (
	"log"
	"net/http"
	"os"
)

//go:generate go run ../genwasm/main.go

func main() {
	log.Print("Starting server ...")
	http.HandleFunc("/", game)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func game(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/ld-50.wasm":
		http.ServeFile(w, r, "ld-50.wasm")
	case "/wasm_exec.js":
		http.ServeFile(w, r, "wasm_exec.js")
	default:
		w.Write([]byte(html))
	}
}

var html = `
<!DOCTYPE html>
<head>
	<title>LD-50 | Droid Battles</title>
</head>
<body>
	<h1>Loading ...</h1>
	<script src="wasm_exec.js"></script>
	<script>
	// Polyfill
	if (!WebAssembly.instantiateStreaming) {
	WebAssembly.instantiateStreaming = async (resp, importObject) => {
		const source = await (await resp).arrayBuffer();
		return await WebAssembly.instantiate(source, importObject);
	};
	}
	const go = new Go();
	WebAssembly.instantiateStreaming(fetch("ld-50.wasm"), go.importObject).then(result => {
	go.run(result.instance);
	});
	</script>
</body>
`
