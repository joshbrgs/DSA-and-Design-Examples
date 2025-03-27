package command

import "fmt"

// Command: This interface defines the method that concrete commands must implement. In our example, the Execute method is used to trigger the command.
// Concrete Commands: These are the specific commands that encapsulate actions. In our example, LightOnCommand and LightOffCommand are concrete commands that turn the light on and off, respectively.
// Invoker: The Button struct acts as an invoker that holds a command and triggers it when needed. The Press method executes the command.
// Client Code: The main function demonstrates how to create concrete commands, set them on the invoker, and trigger them using the invoker.

// Command interface
type Command interface {
	Execute()
}

// Concrete Command for turning on the light
type LightOnCommand struct{}

func (c *LightOnCommand) Execute() {
	fmt.Println("The light is turned on.")
}

// Concrete Command for turning off the light
type LightOffCommand struct{}

func (c *LightOffCommand) Execute() {
	fmt.Println("The light is turned off.")
}

// Invoker
type Button struct {
	command Command
}

// NewButton creates a new Button
func NewButton(command Command) *Button {
	return &Button{command: command}
}

// SetCommand method for Button
func (b *Button) SetCommand(command Command) {
	b.command = command
}

// Press method for Button which triggers the command
func (b *Button) Press() {
	b.command.Execute()
}

func main() {
	lightOn := &LightOnCommand{}
	lightOff := &LightOffCommand{}

	button := NewButton(lightOn)
	button.Press() // Output: The light is turned on.

	button.SetCommand(lightOff)
	button.Press() // Output: The light is turned off.
}
