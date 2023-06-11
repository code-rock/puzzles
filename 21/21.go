// Реализовать паттерн «адаптер» на любом примере.
package main

// Adaptee
type Adaptee struct {
}

func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

// Interface
type Target interface {
	Request() string
}

// Adapter
type Adapter struct {
	*Adaptee
}

func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

func main() {}
