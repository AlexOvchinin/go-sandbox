package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func addAndSub(x, y int) (sum, substraction int) {
	sum = add(x, y)
	substraction = sub(x, y)
	return
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%v > %v\n", v, lim)
	}

	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 30; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%g on the %d iteration\n", z, i)
	}

	return z
}

const E = 2.7343

func main() {
	sqrt := math.Sqrt(float64(rand.Intn(100)))
	fmt.Printf("Now you have %v problems.\n", sqrt)
	fmt.Printf("Now you have %v problems.\n", uint(sqrt))
	var a, b, c = 1, 3, "INIT"
	a, b = addAndSub(12, 23)
	fmt.Printf("Sum is %v\n", a)
	fmt.Printf("Minus is %v\n", b)
	fmt.Printf("Initialized value is %v\n", c)

	var i int
	var f float64
	var k bool
	var s string
	fmt.Printf("Zero values: %v %v %v %q\n", i, f, k, s)

	fmt.Printf("My e is %v\n", E)
	fmt.Printf("Math e is %v\n", math.E)

	multiplication := 1
	for i := 1; i <= 10; i++ {
		multiplication = mul(multiplication, i)
	}

	fmt.Printf("Multiplication = %v\n", multiplication)

	multiplication2 := 1
	var multiplier = 1
	for multiplication2 <= 100000 {
		multiplication2 = mul(multiplication2, multiplier)
		multiplier += 2
	}

	fmt.Printf("Mulltiplication2 = %v\n", multiplication2)

	if multiplication > multiplication2 {
		fmt.Println("First is bigger")
	} else {
		fmt.Println("Second is bigger")
	}

	fmt.Printf("2^10 with 2000 limit = %v, \n", pow(2, 10, 2000))
	fmt.Printf("2^10 with 1000 limit = %v, \n", pow(2, 10, 1000))

	fmt.Printf("Sqrt of %d is %g\n", 123123, Sqrt(123123))

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Printf("%v\n", (today + 7).String())
	switch time.Saturday {
	case today + 6:
		fmt.Println("Today.")
	case today + 7:
		fmt.Println("Tomorrow.")
	case today + 8:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	defer fmt.Printf("world!\n")

	fmt.Printf("Hello ")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
