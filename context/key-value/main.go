package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx, "Hotel")
}

// por convenção o context vem como primeiro parametro
func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Println(token)
}