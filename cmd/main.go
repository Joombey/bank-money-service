package main

import (
	"farukh.go/money/di"
	"farukh.go/money/http"
)

func main() {
	di.Init()
	http.Run()
}