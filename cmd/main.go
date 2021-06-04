package main

import (
	"fmt"

	"github.com/bana118/go-container-template/pkg/example"
)

func main() {
	output := example.Echo("Hello, world!")
	fmt.Println(output)
}
