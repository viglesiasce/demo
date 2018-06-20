package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Host: %s\n", os.Getenv("HOSTNAME"))
	fmt.Printf("Received request from: %s\n", r.Header.Get("X-Forwarded-For"))
}

func main() {
	go func() {
		for {
			fmt.Printf("CPU Limit: %s\n", os.Getenv("MY_CPU_LIMIT"))
			time.Sleep(time.Second * 5)
		}
	}()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
