package main

import "fmt"

func add(num1 int, num2 int) {
	sum := num1 + num2

	fmt.Println(sum)
}

func mul(num1 int, num2 int) int {
	mul := num1 * num2

	return mul
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2

	return sum, mul
}

func printSometing() {
	fmt.Println("Education must be free")
}
func sayHello(name string) {
	fmt.Println("Welcome to the golang course, ", name)
}
func main() {
	printSometing()
	sayHello("Shuvo")
}

/*
	func main() {
		// fmt.Println("Hello World!")

		// data types
		/*
			int
			float32
			bool
			string
*/

// variable declaration
/*
	var a int = 10
	var a = 10
	a := 10
	a := "Hello World!"
	a := true
	const a = 100
	fmt.Println(a)
*/

// conditional statement
// > , >= , < , <= , ==

/*
	age := 18

	if age > 18 {
		fmt.Println("You are eligible to be married")
	} else if age < 18 {
		fmt.Println("You are not eligible to to married")
	} else {
		fmt.Println("You are just a teenager, not eligible to be married")
	}
*/

// and => && or => || not => !
/*
	age := 18
	sex := "male"
	age = 20
	if age == 20 && sex == "male" {
		fmt.Println("You are to be married")
	}
*/

// switch case
/*
	a := 6
	switch a {
	case 1:
		fmt.Println("a is less than 2")
	case 2:
		fmt.Println("a is 2")
	default:
		fmt.Println("a is greater than 2")
	}
*/
/*
		a := 10
		b := 20

		// sum := a + b

		// fmt.Println(sum)
		// add(a, b)

		// multiply := mul(a, b)

		// fmt.Println(multiply)

		p, q := getNumbers(a, b)

		fmt.Println(p, q)

	}

*/
