package scraper

type DataUnit struct {
	name      string // logic name of the data
	dataType  string // type of the data
	data      []string
	selectors map[string][]string // map of data selectors ( for HTML parser )
}
