// The folder's name
package messages

import "fmt"

const message = "Hello from my constant"

// The function will be Uppercase for public access
func Hello() {
	fmt.Println("Hello from messages package")
}

// Public
func Print() {
	fmt.Println(message)
	private()
}

// Private Function
func private() {
	fmt.Println("Hello from private function")

}
