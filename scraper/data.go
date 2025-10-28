package scraper

type DataUnit struct {
	name      string              // logic name of the data
	dataType  string              // type of the data
	data      any                 // will be a slice of the data
	selectors map[string][]string // map of data selectors ( for HTML parser )
}
