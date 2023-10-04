package main

import (
	"context"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	src := client.Host().Directory(".")
	ref := client.
		Container().
		From("golang:1.21-alpine3.18").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src")

	ref.WithExec([]string{"go", "test", "-v"}).Stdout(ctx)
}
