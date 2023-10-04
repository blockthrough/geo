package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		fmt.Printf("dagger.Connect:%s\n", err)
		os.Exit(1)
	}
	src := client.Host().Directory(".")
	ref := client.
		Container().
		From("golang:1.21-alpine3.18").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"go", "test", "-v", "-coverprofile=cov.out", "-coverpkg=./"})

	_, err = ref.Stdout(ctx)
	if err != nil {
		fmt.Printf("go test :%s\n", err)
		os.Exit(1)
	}

	ok, err := ref.File("./cov.out").Export(ctx, "./cov.out")
	if err != nil {
		fmt.Printf("export coverage file :%s\n", err)
		os.Exit(1)
	}

	if !ok {
		fmt.Printf("can not find the coverage file :%s\n", err)
		os.Exit(1)
	}

	// contents, err := ref.WithExec([]string{"go", "test", "-v", "-coverprofile=cov.out", "-coverpkg=./"}).Stdout(ctx)

	// if err != nil {
	// 	fmt.Printf("go test :%s\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("contents:%s\n", contents)
}
