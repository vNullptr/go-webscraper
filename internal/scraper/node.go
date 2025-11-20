package scraper

// no idea if i should make structs for attribute and type
// instead of using a string map
type HTMLAttribute struct {
}

type HTMLType struct {
}

type HTMLNodeWpr struct {
	Data string
	Attr map[string]string
}
