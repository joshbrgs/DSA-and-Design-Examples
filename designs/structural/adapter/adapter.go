package adapter

import "fmt"

// Let's break down the key components of the Adapter Pattern:
//
// Adapter: This is the main component that adapts the existing interface to the expected interface. It contains a reference to the existing interface and implements the expected interface. In our example, PrinterAdapter is the adapter that adapts LegacyPrinter to ModernPrinter.

// Existing Interface: This is the interface that needs to be adapted. In our example, LegacyPrinter is the existing interface.

// Expected Interface: This is the interface that the client expects to use. In our example, ModernPrinter is the expected interface.

// Client: The client is the code that uses the expected interface. In our example, the main function is the client.

// Existing interface that needs to be adapted
type LegacyPrinter interface {
	Print(s string) string
}

// Existing implementation of the LegacyPrinter interface
type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) string {
	newMsg := fmt.Sprintf("Legacy Printer: %s", s)
	fmt.Println(newMsg)
	return newMsg
}

// New interface that the client expects to use
type ModernPrinter interface {
	PrintStored() string
}

// Adapter which adapts LegacyPrinter to ModernPrinter
type PrinterAdapter struct {
	legacyPrinter LegacyPrinter
	msg           string
}

func (p *PrinterAdapter) PrintStored() string {
	// Check if the legacy printer is null (nil) and create a new instance if it is
	if p.legacyPrinter == nil {
		p.legacyPrinter = &MyLegacyPrinter{}
	}
	return p.legacyPrinter.Print(p.msg)
}

func main() {
	legacyPrinter := &MyLegacyPrinter{}
	adapter := &PrinterAdapter{
		legacyPrinter: legacyPrinter,
		msg:           "Hello, World!",
	}

	adapter.PrintStored() // Output: Legacy Printer: Hello, World!
}
