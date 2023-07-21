package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Abser interface {
	Abs() float64
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())
	Scale(&v, 10)
	fmt.Println(Abs(v))
	p := &v
	p.Scale(10)
	fmt.Println(p.Abs())

	var a Abser
	a = f
	a = &v
	fmt.Println(a.Abs())

	var i I

	i = F{math.Pi, "Pi"}
	descrbe(i)
	i.M()

	var t *T
	i = t
	descrbe(i)
	i.M()

	i = &T{"hello"}
	i.M()

	// var i2 I
	// descrbe(i2)
	// i2.M()

	var i3 interface{}

	i3 = 42
	describe(i3)

	i3 = "hello"
	describe(i3)

	typeAssertions()

	do(21)
	do("hello")
	do(true)

	arthur := Person{"Arthur Dent", 42}
	zaphod := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(arthur, zaphod)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

	r := strings.NewReader("Hello, Reader, what the actual")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	var err error
	var n int
	for ; err != io.EOF; n, err = r.Read(b) {
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
	}

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rot := rot13Reader{s}
	io.Copy(os.Stdout, &rot)
	fmt.Println()

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x * y % 256), uint8(x / y % 256), uint8((x + y) / 2 % 256), math.MaxUint8}
}

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	bInt := make([]byte, len(b))
	nInt, errInt := r.r.Read(bInt)
	if errInt != nil {
		return 0, errInt
	}
	for i := 0; i < nInt; i++ {
		if bInt[i] < 'A' || bInt[i] > 'z' {
			b[i] = ' '
		} else {
			if bInt[i] >= 'A' && bInt[i] <= 'M' || bInt[i] >= 'a' && bInt[i] <= 'm' {
				b[i] = bInt[i] + 13
			} else {
				b[i] = bInt[i] - 13
			}
		}
	}
	return nInt, nil
}

type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, err error) {
	bCapped := b[:cap(b)]
	for i := 0; i < len(bCapped); i++ {
		bCapped[i] = 'A'
	}
	return len(bCapped), nil
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number %v", float64(e))
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"It didn't work",
	}
}

type IPAddr [4]byte

func (addr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", addr[0], addr[1], addr[2], addr[3])
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about the type %T\n", v)
	}
}

func typeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64)
	// fmt.Println(f)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func descrbe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type F struct {
	f float64
	s string
}

func (f F) M() {
	fmt.Println(f.f)
}

type I interface {
	M()
}

type T struct {
	S string
}
