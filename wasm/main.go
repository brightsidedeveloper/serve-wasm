package main

import (
	"wasm/app/drivers"

	"github.com/brightsidedeveloper/goat"
)

func main() {
	done := make(chan struct{}, 0)

	goat.Log("Hi from WASM!")

	drivers.RouterProvider()

	<-done
}
