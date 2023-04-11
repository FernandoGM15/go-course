package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServer(server string, channel chan string) {
	_, error := http.Get(server)
	if error != nil {
		// fmt.Printf("%s is not available\n", server)
		channel <- server + " is not available"
	} else {
		// fmt.Printf("%s It's working\n", server)
		channel <- server + " Is working"
	}
}

func main() {

	// Time from start
	start := time.Now()

	//Making a channel for the concurrency process
	channel := make(chan string)

	servers := []string{
		"https://www.oregoom.com",
		"https://www.udemy.com",
		"https://www.facebook.com",
		"https://www.google.com",
	}

	// With concurrency
	for _, server := range servers {
		go checkServer(server, channel)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}

	// Difference between start and end time
	time := time.Since(start)

	fmt.Println("Execution time ", time)
}
