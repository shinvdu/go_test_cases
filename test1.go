package main

import "fmt"
import "time"

func main() {
	// generatedPassword := make(chan []byte, 100)
	generatedPassword := make(chan []byte)
	correctPassword := make(chan []byte)
	defer close(generatedPassword)
	defer close(correctPassword)

	go passwordIncrement(generatedPassword)

	go checkPassword(generatedPassword, correctPassword)
	pass := <-correctPassword
	fmt.Println(string(pass))
}

func checkPassword(input <-chan []byte, output chan<- []byte) {
	for {
		p := <-input
		//Introduce lengthy operation here
		time.Sleep(time.Second)
		// fmt.Println("Checking p:", string(p))
		// time.Sleep(1 * 1e9) 
		if performSomeCheckingOperation(p) {
			fmt.Println("Checking p:", string(p))
			output <- p
		}
	}
}

func passwordIncrement(out chan<- []byte) {
	p := []byte("abc")

	for {
		p = generate(p)
		out <- p
	}
}

func generate(p []byte)(t []byte) {

	p = append(p, []byte("d")...)
	return p
	// return []byte("abcddd")
}

func performSomeCheckingOperation(p []byte)(a bool){
	fmt.Println("Checking p:", string(p))

	if string(p) == "abcdddddd" {
		return true
	}else{
		return false
	}
}
