package main

import (
	"fmt"
	"sync"
	"webscraper/scraper"
)


var wg sync.WaitGroup
func test(){

	s := scraper.NewScraper()
	wg.Add(1)
	s.DlUrl("https://www.example.com", "GET", nil)
	fmt.Println(string(s.GetUnparsedHTML()))
	wg.Done()

}

func main() {

	// testing goroutine ( for concurrency implementation )
	go test()
	fmt.Println("test")
	wg.Wait()

}
