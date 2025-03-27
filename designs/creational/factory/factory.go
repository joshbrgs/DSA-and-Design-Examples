package factory

import "fmt"

// Key Components of the Factory Pattern
// Product: The interface or abstract class that defines the type of objects the factory method can create. In our example, Document is the product interface.

// Concrete Product: The concrete classes that implement the Product interface. In our example, WordDocument and ExcelDocument are concrete products.

// Creator: The interface or abstract class that declares the factory method for creating objects. In our example, DocumentFactory is the creator interface.

// Concrete Creator: The concrete classes that implement the Creator interface and define the factory method to create objects. In our example, WordDocumentFactory and ExcelDocumentFactory are concrete creators.

// Client: The code that uses the factory method to create objects. In our example, the main function acts as the client.

// Document is an interface that all concrete documents will implement
type Document interface {
	PrintDocument()
}

// DocumentFactory is an interface defining the factory method
type DocumentFactory interface {
	CreateDocument() Document
}

// WordDocument is a concrete type implementing Document interface
type WordDocument struct{}

func (w WordDocument) PrintDocument() {
	fmt.Println("This is a Word document.")
}

// WordDocumentFactory is a factory for creating Word documents
type WordDocumentFactory struct{}

func (w WordDocumentFactory) CreateDocument() Document {
	return WordDocument{}
}

// ExcelDocument is a concrete type implementing Document interface
type ExcelDocument struct{}

func (e ExcelDocument) PrintDocument() {
	fmt.Println("This is an Excel document.")
}

// ExcelDocumentFactory is a factory for creating Excel documents
type ExcelDocumentFactory struct{}

func (e ExcelDocumentFactory) CreateDocument() Document {
	return ExcelDocument{}
}

func factoryMain() {
	var docFactory DocumentFactory

	// Use WordDocumentFactory to create a Word document
	docFactory = WordDocumentFactory{}
	wordDoc := docFactory.CreateDocument()
	wordDoc.PrintDocument() // Output: This is a Word document.

	// Use ExcelDocumentFactory to create an Excel document
	docFactory = ExcelDocumentFactory{}
	excelDoc := docFactory.CreateDocument()
	excelDoc.PrintDocument() // Output: This is an Excel document.
}
