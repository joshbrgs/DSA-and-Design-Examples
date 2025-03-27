package observer

import "fmt"

type Observer interface {
	Update(string)
}

type ConcreteObserver struct {
	ID   int
	Name string
}

func (co *ConcreteObserver) Update(message string) {
	fmt.Printf("Subscriber %s (ID: %d) received weather update: %s\n", co.Name, co.ID, message)
}

type Subject interface {
	RegisterObserver(Observer)
	UnregisterObserver(Observer)
	NotifyObservers(string)
}

type ConcreteSubject struct {
	observers []Observer
}

func (cs *ConcreteSubject) RegisterObserver(o Observer) {
	cs.observers = append(cs.observers, o)
}

func (cs *ConcreteSubject) UnregisterObserver(o Observer) {
	var indexToRemove int
	for index, observer := range cs.observers {
		if observer == o {
			indexToRemove = index
			break
		}
	}
	cs.observers = append(cs.observers[:indexToRemove], cs.observers[indexToRemove+1:]...)
}

func (cs *ConcreteSubject) NotifyObservers(message string) {
	for _, observer := range cs.observers {
		observer.Update(message)
	}
}

func main() {
	subscriber1 := &ConcreteObserver{ID: 1, Name: "WeatherApp"}
	subscriber2 := &ConcreteObserver{ID: 2, Name: "NewsWebsite"}

	weatherStation := &ConcreteSubject{}

	weatherStation.RegisterObserver(subscriber1)
	weatherStation.RegisterObserver(subscriber2)

	weatherStation.NotifyObservers("Sunny with a high of 25°C") // Both observers receive the message and print it to the console

	weatherStation.UnregisterObserver(subscriber1)

	weatherStation.NotifyObservers("Rainy with a chance of thunderstorms") // Only subscriber2 receives the message, as subscriber1 was unregistered
}

// We define the Observer interface with an Update method that observers must implement.
// The ConcreteObserver struct implements the Observer interface and prints the received message when the Update method is called.
// The Subject interface defines methods to register, unregister, and notify observers.
// The ConcreteSubject struct maintains a list of observers and implements the Subject interface methods.
// In the main function, we create two observers (subscriber1 and subscriber2) and a subject (weatherStation).
// We register the observers with the subject, notify them of a state change, and then unregister one of the observers before sending another notification.
// When the subject notifies the observers, they receive the message and print it to the console with their respective names and IDs.
