package main

import (
	"fmt"
	"sync"
)

// - Insert into the "output" var all Nike + Adidas + Puma product IDs that start with the digit "1"
// - Do not change the the 3 Get functions in any way
// - Utilize go routines (i.e. you will want to be calling the Get functions and inserting into "output" concurrently)
var wg sync.WaitGroup

func main() {
	// utilize the max num of cores available
	// runtime.GOMAXPROCS(runtime.NumCPU())

	var output []string
	c := make(chan string, 100)
	// c := make(chan string)
	//  short on time today, but might also try an unbuffered example latter.

	wg.Add(1)
	go getshoes02(c)
	wg.Wait()

	for item := range c {
		fmt.Println("+++!!  received - ", item)
		output = append(output, item)
	}
	fmt.Println(output)
}

func getshoes02(c chan string) {
	defer wg.Done()
	defer close(c)

	nn := GetNikeProductIDs()
	aa := GetAdidasProductIDs()
	pp := GetPumaProductIDs()

	for _, v := range nn {
		if testReturn(v) {
			fmt.Println("N Sending v in channel", v)
			c <- v
		} else {
			// fmt.Println("N not sending v in channel", v)
		}
	}

	for _, v := range aa {
		// fmt.Println("shoe: ", shoe, " = ", v, " - ", testReturn(v))
		if testReturn(v) {
			fmt.Println("N Sending v in channel", v)
			c <- v
		} else {
			// fmt.Println("N not sending v in channel", v)
		}
	}

	for _, v := range pp {
		// fmt.Println("shoe: ", shoe, " = ", v, " - ", testReturn(v))
		if testReturn(v) {
			fmt.Println("N Sending v in channel", v)
			c <- v
		} else {
			// fmt.Println("N not sending v in channel", v)
		}
	}
}

//================================================
func GetNikeProductIDs() (out []string) {
	for i := 0; i < 100; i += 10 {
		out = append(out, fmt.Sprintf("%d_%s", i, "nike"))
	}
	return out
}

func GetAdidasProductIDs() (out []string) {
	for i := 0; i < 100; i += 5 {
		out = append(out, fmt.Sprintf("%d_%s", i, "adidas"))
	}
	return out
}

func GetPumaProductIDs() (out []string) {
	for i := 0; i < 100; i += 2 {
		out = append(out, fmt.Sprintf("%d_%s", i, "puma"))
	}
	return out
}

//====================================

func testReturn(s string) bool {
	if s[0:1] == "1" {
		return true
	} else {
		return false
	}
}
