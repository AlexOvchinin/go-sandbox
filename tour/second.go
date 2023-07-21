package main

import (
	"fmt"
	"math"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{3, 4}
	v3 = Vertex{}
	p  = &Vertex{5, 6}
)

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println(p)
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
	fmt.Printf("%+v\n", v)

	vp := &v
	vp.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World!"
	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:5]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"Ringo",
		"George",
	}
	fmt.Println(names)

	aa := names[0:2]
	b := names[1:3]
	fmt.Println(aa, b)

	b[0] = "XXX"
	fmt.Println(aa, b)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	structs := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(structs)

	var slice = primes[:]
	slice = slice[1:4]
	fmt.Println(slice)

	slice = slice[:2]
	fmt.Println(slice)

	slice = slice[1:]
	fmt.Println(slice)

	slice = primes[:0]
	printSlice(slice)

	slice = slice[:4]
	printSlice(slice)

	slice = slice[:2]
	printSlice(slice)

	slice = slice[:4]
	printSlice(slice)

	slice = slice[2:]
	printSlice(slice)

	var nilSlice []int
	fmt.Println(s, len(s), cap(s))
	if nilSlice == nil {
		fmt.Println("nil!")
	}

	makeSlice()
	sliceOfSlices()
	appendSlice()
	rangeFor()
	rangeSkipped()
	rangeSkippedArray()
	fmt.Println(Pic(10, 10))
	createMap()
	mapLiterals()
	mutatingMaps()
	fmt.Println(WordCount("123 321 123"))
	functionValues()
	closures()
	fibCaller()
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first, second := 1, 1
	return func() int {
		result := first + second
		first = second
		second = result
		return result
	}
}

func fibCaller() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func closures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func WordCount(s string) map[string]int {
	var result map[string]int = make(map[string]int)
	for _, value := range strings.Fields(s) {
		elem, ok := result[value]
		if ok {
			result[value] = elem + 1
		} else {
			result[value] = 1
		}
	}
	return result
}

func mutatingMaps() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value: ", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func mapLiterals() {
	var m = map[string]NewVertex{
		"Bell Labs": {36, -123.2},
		"Google":    {12.213, 12312.123},
	}
	fmt.Println(m)
}

type NewVertex struct {
	Lat, Long float64
}

func createMap() {
	var m map[string]NewVertex

	m = make(map[string]NewVertex)
	m["Bell Labs"] = NewVertex{40.8, -74.36}
	fmt.Println(m["Bell Labs"])
}

func Pic(dx, dy int) [][]uint8 {
	var result = make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		result[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			result[i][j] = uint8(i ^ j)
		}
	}
	return result
}

func rangeSkippedArray() {
	pow := [2]int{0, 0}
	for i := range pow {
		pow[i] = 1 << i
	}

	for _, value := range pow {
		fmt.Printf("%v\n", value)
	}
}

func rangeSkipped() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << i
	}

	for _, value := range pow {
		fmt.Printf("%v\n", value)
	}
}

func rangeFor() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func appendSlice() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	s = append(s, 7, 8, 9)
	printSlice(s)
}

func sliceOfSlices() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func makeSlice() {
	a := make([]int, 5)
	printSliceWithString("a", a)

	b := make([]int, 0, 5)
	printSliceWithString("b", b)
	// b[0] = 1
	// b[2] = 3
	// b[4] = 5

	c := b[:2]
	printSliceWithString("c", c)

	d := c[2:5]
	printSliceWithString("d", d)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceWithString(s string, slice []int) {
	fmt.Printf("%s len = %d cap = %d %v\n", s, len(slice), cap(slice), slice)
}
