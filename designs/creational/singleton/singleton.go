package singleton

import (
	"fmt"
	"sync"
)

// Use Cases
// Configuration Management: A singleton can store global configuration settings for an application, ensuring that all parts of the program access a consistent and unified configuration.

// Logging: Centralize logging to a single instance to ensure that all logs are collected in a single place, making it easier to manage and analyze.

// Thread Pool Management: Prevent the creation of multiple thread pools and ensure that all tasks share a single thread pool resource.

// Caching: Implement a caching mechanism that is accessible globally by all parts of an application to improve performance and resource usage.

// Database Connections: Manage connections to a database to ensure that only one connection pool is used throughout the application, optimizing resource allocation and performance.

// File System Access: Control access to files to ensure that only one instance handles read and write operations, preventing data corruption or inconsistencies.

type Logger struct{}

var instance *Logger
var once sync.Once

// Function to get the singleton instance
func GetInstance() *Logger {
	once.Do(func() {
		instance = &Logger{}
	})
	return instance
}

func (l *Logger) Log(message string) {
	fmt.Println(message)
}

func main() {
	logger1 := GetInstance()
	logger2 := GetInstance()

	if logger1 == logger2 {
		fmt.Println("Both instances are the same.") // This block will be executed since singleton is the same instance
	} else {
		fmt.Println("Instances are different.")
	}

	logger1.Log("Logging a message")       // Output: Logging a message
	logger2.Log("Logging another message") // Output: Logging another message
}

// Go's concurrency model, centered around goroutines and channels, integrates well with the Singleton Pattern when thread safety is ensured using primitives like sync.Once. This makes it possible to implement Singletons in a way that is consistent with Go's design philosophy of providing powerful concurrency constructs while keeping the code simple and clear.
