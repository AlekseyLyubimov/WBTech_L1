package main

import "strings"

func main() {
	adaptee := Adaptee{}
	adapter := Adapter{Adaptee: &adaptee}
	println("adapted message:")
	println(adapter.Request())
}

type Adaptee struct {
}

func (a *Adaptee) SpecificRequest() string {
	return "message to adapt"
}

type Adapter struct {
	*Adaptee
}

func (a *Adapter) Request() string {
	return strings.ToUpper(a.SpecificRequest())
}
