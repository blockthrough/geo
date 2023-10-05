package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	covFile := flag.String("cov_file", "cov.out", "specify the cov file output")
	flag.Parse()

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
		WithExec([]string{"go", "test", "-v", fmt.Sprintf("-coverprofile=%s", *covFile), "-coverpkg=./"})

	_, err = ref.Stdout(ctx)
	if err != nil {
		fmt.Printf("go test :%s\n", err)
		os.Exit(1)
	}

	// retrieve the coverage file
	ok, err := ref.File(*covFile).Export(ctx, *covFile)
	if err != nil {
		fmt.Printf("export coverage file :%s\n", err)
		os.Exit(1)
	}

	if !ok {
		fmt.Printf("can not find the coverage file :%s\n", err)
		os.Exit(1)
	}

}
