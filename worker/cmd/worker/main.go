package main

import (
	// "fmt"
	"go-async-training-worker/internal/jobs"
)

func main() {
	channel := make(chan [][]string)

	go jobs.Parse("/app/uploads/customers-1000.csv", channel)
	jobs.Validate(channel)
	// fmt.Println(channel)
}
