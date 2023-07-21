package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// go say("World")
	// say("Hello")

	// s := []int{7, 2, 8, -9, 4, 0}

	// c := make(chan int, 2)
	// go sum("first", s[:len(s)/2], c)
	// go sum("second", s[len(s)/2:], c)
	// x, y := <-c, <-c

	// fmt.Println(x, y, x+y)

	// ch := make(chan int, 2)
	// ch <- 1
	// ch <- 2
	// // ch <- 3
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// rangeAndClose()
	// defaultSelection()

	// c := SafeCounter{v: make(map[string]int)}
	// for i := 0; i < 1000; i++ {
	// go c.Inc("somekey")
	// }

	// time.Sleep(time.Second)
	// fmt.Println(c.Value("somekey"))
	go Crawl("https://golang.org/", 4, fetcher)

	time.Sleep(time.Second)
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

var crawledUrls = make(map[string]int)
var crawledUrlsMutex = sync.Mutex{}
var wg = sync.WaitGroup{}

func Crawl(url string, depth int, fetcher Fetcher) {
	crawledUrlsMutex.Lock()
	if _, ok := crawledUrls[url]; ok {
		crawledUrlsMutex.Unlock()
		return
	}
	crawledUrls[url] = 1
	crawledUrlsMutex.Unlock()

	wg.Add(1)
	defer wg.Done()

	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func Walk(t *Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		v1 := <-ch1
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func defaultSelection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func selectFunc() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 15; i++ {
			var cVal = <-c
			fmt.Println(cVal)
		}
		quit <- 0
	}()
	fibonacciSelect(c, quit)
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case q := <-quit:
			fmt.Println(q)
			fmt.Println("quit")
			return
		}
	}
}

func rangeAndClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func sum(name string, s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		fmt.Printf("Goroutine: %v, Current sum %v\n", name, sum)
	}
	c <- sum //send sum to c
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
