package cor

import "fmt"

// Handler Interface: This interface defines the common methods that all handlers must implement. In our example, the Handle method processes the request, and the SetNext method sets the next handler in the chain.
// BaseHandler: This is a base class that implements the Handler interface. It contains the logic to pass the request
// Concrete Handlers: These are specific handlers that extend the BaseHandler and implement the Handle method. They decide whether to handle the request or pass it to the next handler in the chain. In our example, we have FrontDeskHandler, ManagerHandler, and DirectorHandler.
// Client Code: This is where you set up the chain of handlers and initiate the request. In our example, we create instances of the handlers, link them together, and pass the request to the first handler in the chain.

// Let's now understand the flow of the request in our example:

// The Request object is passed to the FrontDeskHandler. If the request level matches the handler's level, it handles the request; otherwise, it passes the request to the next handler in the chain.
// The ManagerHandler receives the request and checks if it can handle it. If not, it passes the request to the DirectorHandler.
// The DirectorHandler processes the request if it matches its level.
// If none of the handlers can handle the request, it remains unhandled.
// The client code initiates the request by calling the Handle method of the first handler in the chain.

type Request struct {
	message string
	level   int
}

type Handler interface {
	Handle(request Request) bool
	SetNext(handler Handler)
}

// Base Handler
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) Handle(request Request) bool {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return false
}

// Front Desk Handler
type FrontDeskHandler struct {
	BaseHandler
}

func (h *FrontDeskHandler) Handle(request Request) bool {
	if request.level == 1 {
		fmt.Println("Request handled by Front Desk")
		return true
	}
	return h.BaseHandler.Handle(request)
}

// Manager Handler
type ManagerHandler struct {
	BaseHandler
}

func (h *ManagerHandler) Handle(request Request) bool {
	if request.level == 2 {
		fmt.Println("Request handled by Manager")
		return true
	}
	return h.BaseHandler.Handle(request)
}

// Director Handler
type DirectorHandler struct {
	BaseHandler
}

func (h *DirectorHandler) Handle(request Request) bool {
	if request.level == 3 {
		fmt.Println("Request handled by Director")
		return true
	}
	return h.BaseHandler.Handle(request)
}

func main() {
	frontDesk := &FrontDeskHandler{}
	manager := &ManagerHandler{}
	director := &DirectorHandler{}

	frontDesk.SetNext(manager)
	manager.SetNext(director)

	request := Request{message: "General inquiry", level: 1}
	frontDesk.Handle(request) // Output: Request handled by Front Desk

	request = Request{message: "Complaint", level: 2}
	frontDesk.Handle(request) // Output: Request handled by Manager

	request = Request{message: "High-level issue", level: 3}
	frontDesk.Handle(request) // Output: Request handled by Director
}
