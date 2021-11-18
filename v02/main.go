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
	var output []string
	output = getshoes()
	fmt.Println(output)
}

func getshoes() []string {

	var result []string
	// c := make(chan string)
	c := make(chan string, 1000)

	wg.Add(1)
	go getshoe(c, "N")
	wg.Add(1)
	go getshoe(c, "P")
	wg.Add(1)
	go getshoe(c, "A")
	wg.Wait()

	close(c)
	for v := range c {
		result = append(result, v)
	}
	return result
}

func getshoe(c chan string, shoe string) {
	var nn []string

	switch shoe {
	case "N":
		fmt.Println("0000  starting N")
		nn = GetNikeProductIDs()
	case "A":
		fmt.Println("0000  starting A")
		nn = GetAdidasProductIDs()
	case "P":
		fmt.Println("0000  starting P")
		nn = GetPumaProductIDs()
	}

	for _, v := range nn {
		if testReturn(v) {
			fmt.Println("N Sending v in channel", v)
			c <- v
		} else {
			// fmt.Println("N not sending v in channel", v)
		}
	}
	wg.Done()

	return
}

//--------------------------------------------
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
