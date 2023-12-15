package main

import (
	"farukh.go/money/db"
	"farukh.go/money/http"
)

func main() {
	// if path := os.Getenv("CONFIG_PATH"); path == "" {
	// 	in.Init(cts.LocalConfigPath)
	// } else {
	// 	in.Init(path)
	// }
	db.Init()
	http.Run()
}
