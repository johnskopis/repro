package main

import (

	_ "github.com/mattn/go-sqlite3"
	"context"
)


func main() {
	ctx := context.Background()
	// never reaches
	fmt.Printf("ok")
}
