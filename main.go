package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("ENV: %s\n", os.Getenv("TEST_SECRET"))
}
