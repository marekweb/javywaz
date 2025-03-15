package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/marekweb/javywaz/pkg/executor"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Instantiate the JavyExecutor with the path to your WASM binary.
	executor, err := executor.NewJavyExecutor(ctx, "example.wasm")
	if err != nil {
		log.Fatal(err)
	}
	defer executor.Close(ctx)

	// Provide the JSON input as a string.
	jsonInput := `{"n": 1337, "bar": "hello"}`
	result, err := executor.Execute(ctx, jsonInput)
	if err != nil {
		log.Fatal(err)
	}

	// Print the output from stdout.
	fmt.Println("Output:")
	fmt.Println(result.Stdout)
}
