package main

import (
	"webscrapper/internal/fetcher"
	"fmt"
)


func main() {

	fetcher := fetcher.HttpClient{
		Url: "http://example.com",
		Method: "GET",
	}
	var result, _ = fetcher.Fetch()
	fmt.Println(result)

}