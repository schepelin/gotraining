// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to read a stack trace.
package main

func main() {
	example(make([]string, 2, 4), "hello", 10)
}

//go:noinline
func example(slice []string, str string, i int) {
	panic("Want stack trace")
}

/*
	panic: Want stack trace

	goroutine 1 [running]:
	main.example(0xc000042748, 0x2, 0x4, 0x106abae, 0x5, 0xa)
		stack_trace/example1/example1.go:13 +0x39  <--- hex offset of instruction ()
	main.main()
		stack_trace/example1/example1.go:8 +0x72

	--------------------------------------------------------------------------------

	// Declaration
	main.example(slice []string, str string, i int)

	// Call
	main.example(0xc000042748, 0x2, 0x4, 0x106abae, 0x5, 0xa)

	// Stack trace
	main.example(0xc000042748, 0x2, 0x4, 0x106abae, 0x5, 0xa)

	// Values
	Slice Value:   0xc000042748, 0x2, 0x4
	String Value:  0x106abae, 0x5
	Integer Value: 0xa
*/

// Note: https://go-review.googlesource.com/c/go/+/109918
