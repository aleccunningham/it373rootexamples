package main

//go:noinline
func addThree() int32 {
	var a int32
	a = 1000
	return a + 3
}

func main() { addThree() }
