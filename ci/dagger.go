package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"dagger.io/dagger"
)

const DockerImage string = "golang:1.21-alpine3.18"

func main() {
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		log.Println(err)
		os.Exit(1)
		return
	}

	defer client.Close()
	src := client.Host().Directory(".")

	golang := client.Container().
		From(DockerImage).WithDirectory("/src", src).WithWorkdir("/src")

	unitTest := golang.WithExec([]string{"go", "test", "-v"})

	output, err := unitTest.Stdout(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(output)
}
