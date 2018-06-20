package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		fmt.Printf("Hello world from %s!\n", os.Getenv("HOSTNAME"))
		time.Sleep(time.Second * 1)
	}
}
