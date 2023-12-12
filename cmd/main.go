package main

import (
	"os"

	cts "farukh.go/money/constants"
	"farukh.go/money/db"
	"farukh.go/money/http"
	in "farukh.go/money/internal"
)

func main() {
	if path := os.Getenv("CONFIG_PATH"); path == "" {
		in.Init(cts.LocalConfigPath)
	} else {
		in.Init(path)
	}
	db.Init()
	http.Run()
}
