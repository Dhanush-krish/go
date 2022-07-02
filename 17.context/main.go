package main

import (
	"context"
	"fmt"
)

// context is key-value pair channel that can be accessed across all the layers

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "stored-in-main-function")

	retriveContext(ctx)
}

func retriveContext(ctx context.Context) {
	fmt.Println("accessing value pushed in main function:", ctx.Value("key"))
}
