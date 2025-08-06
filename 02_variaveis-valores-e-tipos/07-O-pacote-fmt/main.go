package main

import "fmt"

func main() {
	fmt.Println("Hello", "World")
	fmt.Print("Hello ", "World") // Outputs: Hello World (no newline)
	name := "Go"
	version := 1.18
	fmt.Printf("Language: %s, Version: %.1f\n", name, version) // Language: Go, Version: 1.8
	/*
		%v - Default format
		%T - Type of the value
		%d - Integer
		%f - Float
		%s - String
		%t - Boolean
		%p - Pointer
	*/
	//Returns a formatted  string without printing it:
	message := fmt.Sprintf("Score: %d", 100)
	fmt.Print(message)

	//fmt.Scan // Reads input from standard input (console)
	var name2 string
	var age int
	fmt.Scan(&name2, &age)

	//Reads input until a newline:
	var input string
	fmt.Scanln(&input)

	//Most functions return the number of bytes written and an error value:
	count, err := fmt.Println("Hello")
	fmt.Print(count)
	if err != nil {
		// Handle error
	}

	//The Fprint Family of Functions in fmt Package

	//Writes values to a specified writer without adding a newline:
	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "Hello", "World") // Writes: HelloWorld

	//Writes values to a specified writer with spaces between them and adds a newline:
	file, _ := os.Create("output.txt")
	fmt.Fprintln(file, "Hello", "World") // Writes: Hello World\n

	//Writes formatted text to a specified writer:
	file, _ := os.Create("output.txt")
	name := "Go"
	fmt.Fprintf(file, "Language: %s, Version: %d\n", name, 1)

	/*
		Common Use Cases
		These functions are particularly useful for:


		Writing logs to files
		Sending formatted data over network connections
		Building strings in buffers (bytes.Buffer)
		Testing output by writing to custom writers
		All Fprint functions return the number of bytes
		written and any error encountered, making error handling straightforward.
	*/
}
