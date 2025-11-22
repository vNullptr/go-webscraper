package scraper

// no idea if i should make structs for attribute and type
// instead of using a string map
type HTMLAttribute struct {
	Name, Val string
}

type HTMLType uint16

type HTMLNodeWpr struct {
	Data string
	Attr HTMLAttribute
	Type HTMLType 
}
