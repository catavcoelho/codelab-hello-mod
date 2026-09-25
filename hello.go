// Package hellomod provides greeting utilities for the Go Module Proxy codelab.
package hellomod

import "fmt"

// Greet returns a friendly greeting for the given name.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// unexportedHelper is only visible inside package hellomod.
func unexportedHelper() string {
	return "internal helper"
}

