package strategy

import (
	"fmt"
	"strings"
)

// Strategy: An interface that defines the method set for algorithms. It allows multiple implementations of the same algorithm. In our example, Strategy defines the Format method for string formatting algorithms.
// Concrete Strategies: Different implementations of the Strategy interface. In our example, UpperCase and LowerCase are concrete strategies that implement the Format method.
// Context: A class that holds a reference to a Strategy object and uses it to execute the algorithm. The context class allows you to switch between different strategies at runtime. In our example, Context holds a reference to a Strategy and uses it to format strings.
// Client: The client code that uses the Context class to execute algorithms. In our example, the main function creates a Context object, sets different strategies, and executes them.

// Strategy defines the method set for algorithms.
type Strategy interface {
	Format(str string) string
}

// UpperCase strategy.
type UpperCase struct{}

func (UpperCase) Format(str string) string {
	return strings.ToUpper(str)
}

// LowerCase strategy.
type LowerCase struct{}

func (LowerCase) Format(str string) string {
	return strings.ToLower(str)
}

// Context holds a reference to a Strategy.
type Context struct {
	strategy Strategy
}

// SetStrategy sets the strategy in the context.
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// Execute uses the set strategy to format the given string.
func (c *Context) Execute(str string) string {
	return c.strategy.Format(str)
}

func main() {
	context := &Context{}

	input := "the quick brown fox jumps over the lazy dog"

	// Set and execute UpperCase strategy.
	context.SetStrategy(UpperCase{})
	result := context.Execute(input)
	fmt.Printf("UpperCase Strategy Result: %s\n", result) // Output: THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG

	// Set and execute LowerCase strategy.
	context.SetStrategy(LowerCase{})
	result = context.Execute(input)
	fmt.Printf("LowerCase Strategy Result: %s\n", result) // Output: the quick brown fox jumps over the lazy dog
}
