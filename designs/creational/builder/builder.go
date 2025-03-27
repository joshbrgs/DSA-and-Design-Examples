package builder

import "fmt"

// A Builder Pattern is useful when you need to construct complex objects step by step, ensuring an easy-to-read and flexible object creation process. It is particularly beneficial when dealing with objects that have many optional parameters.

// This pattern is commonly used in creating configuration objects, constructing SQL queries, and setting up complex API requests.

// Report represents the complex object we want to build
type Report struct {
	Title   string
	Author  string
	Content string
	Footer  string
	PageNum int
}

func (r Report) PrintReport() {
	fmt.Println("---- Report ----")
	fmt.Printf("Title: %s\n", r.Title)
	fmt.Printf("Author: %s\n", r.Author)
	fmt.Printf("Content: %s\n", r.Content)
	fmt.Printf("Footer: %s\n", r.Footer)
	fmt.Printf("Page Numbers: %d\n", r.PageNum)
}

// ReportBuilder helps construct a Report step by step
type ReportBuilder struct {
	report Report
}

// NewReportBuilder creates a new builder instance
func NewReportBuilder() *ReportBuilder {
	return &ReportBuilder{}
}

// SetTitle sets the title of the report
func (b *ReportBuilder) SetTitle(title string) *ReportBuilder {
	b.report.Title = title
	return b
}

// SetAuthor sets the author of the report
func (b *ReportBuilder) SetAuthor(author string) *ReportBuilder {
	b.report.Author = author
	return b
}

// SetContent sets the content of the report
func (b *ReportBuilder) SetContent(content string) *ReportBuilder {
	b.report.Content = content
	return b
}

// SetFooter sets the footer of the report
func (b *ReportBuilder) SetFooter(footer string) *ReportBuilder {
	b.report.Footer = footer
	return b
}

// SetPageNumbers sets the number of pages in the report
func (b *ReportBuilder) SetPageNumbers(pages int) *ReportBuilder {
	b.report.PageNum = pages
	return b
}

// Build returns the final constructed Report
func (b *ReportBuilder) Build() Report {
	return b.report
}

// See how the builder pattern always returns the ReportBuilder, that is the key to the pattern letting you use it like this:

func main() {
	builder := NewReportBuilder()
	report := builder.
		SetTitle("Builder Pattern in Golang").
		SetAuthor("John Doe").
		SetContent("The builder pattern helps create complex objects step by step.").
		SetFooter("Confidential").
		SetPageNumbers(10).
		Build()

	report.PrintReport()
}
