package abstractfactory

import "fmt"

// Encapsulation of Object Creation: Clients don't need to know the concrete implementations.
//
// Easier Scalability: Adding a new document type only requires adding a new method to the factory.
//
// Ensures Consistency: Factories produce a family of related products.

// Document is an interface that all concrete documents will implement
type Document interface {
	PrintDocument()
}

// WordDocument is a concrete type implementing the Document interface
type WordDocument struct{}

func (w WordDocument) PrintDocument() {
	fmt.Println("This is a Word document.")
}

// ExcelDocument is a concrete type implementing the Document interface
type ExcelDocument struct{}

func (e ExcelDocument) PrintDocument() {
	fmt.Println("This is an Excel document.")
}

// PowerPointDocument is a concrete type implementing the Document interface
type PowerPointDocument struct{}

func (p PowerPointDocument) PrintDocument() {
	fmt.Println("This is a PowerPoint document.")
}

// DocumentAbstractFactory is an abstract factory interface that defines methods to create documents
type DocumentAbstractFactory interface {
	CreateWordDocument() Document
	CreateExcelDocument() Document
	CreatePowerPointDocument() Document
}

// OfficeDocumentFactory is a concrete factory that implements DocumentAbstractFactory
type OfficeDocumentFactory struct{}

func (o OfficeDocumentFactory) CreateWordDocument() Document {
	return WordDocument{}
}

func (o OfficeDocumentFactory) CreateExcelDocument() Document {
	return ExcelDocument{}
}

func (o OfficeDocumentFactory) CreatePowerPointDocument() Document {
	return PowerPointDocument{}
}

func main() {
	var docFactory DocumentAbstractFactory

	// Use OfficeDocumentFactory to create different types of documents
	docFactory = OfficeDocumentFactory{}

	wordDoc := docFactory.CreateWordDocument()
	wordDoc.PrintDocument() // Output: This is a Word document.

	excelDoc := docFactory.CreateExcelDocument()
	excelDoc.PrintDocument() // Output: This is an Excel document.

	pptDoc := docFactory.CreatePowerPointDocument()
	pptDoc.PrintDocument() // Output: This is a PowerPoint document.
}
