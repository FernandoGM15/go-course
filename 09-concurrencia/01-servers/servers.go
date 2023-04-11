package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServer(server string) {
	_, error := http.Get(server)
	if error != nil {
		fmt.Printf("%s is not available\n", server)
	} else {
		fmt.Printf("%s It's working\n", server)
	}
}

func main() {

	// Time from start
	start := time.Now()

	servers := []string{
		"https://www.oregoom.com",
		"https://www.udemy.com",
		"https://www.facebook.com",
		"https://www.google.com",
	}

	// Without concurrency
	for _, server := range servers {
		checkServer(server)
	}

	// Difference between start and end time
	time := time.Since(start)

	fmt.Println("Execution time ", time)
}
