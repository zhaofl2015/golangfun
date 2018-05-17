package main

type decorator func()

func (this decorator) invoke() {
	println("before")
	this()
	println("after")
}

func myFunc() {
	println("blah")
}

func main() {
	x := decorator(myFunc)
	x.invoke()
}
