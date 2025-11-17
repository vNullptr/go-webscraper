package scraper

import (
	"fmt"
	"sync"

	"golang.org/x/net/html"
)

type Scraper struct {
	// to be implemented
	// will regroup everything needed for scraping
	// from result of url download to parsing result ect ...

	mu           sync.RWMutex
	unparsedHTML []byte
	targetData   []DataUnit

	// this might stay unused for a little while but will store the root of page
	// will also be copied and clean up into another node tree
	htmlRoot *html.Node
}

// thinking about making timeout tied to the scraper instance
func NewScraper() *Scraper {
	return &Scraper{}
}

// still defaulting to string type haven't done any testing with the interface
func (s *Scraper) DataUnit(name string, dataType string, selectorMap map[string][]string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	du := DataUnit{
		name:      name,
		dataType:  dataType,
		data:      []string{}, 
		selectors: selectorMap,
	}

	s.targetData = append(s.targetData, du)
}


func (s *Scraper) DebugShowData(){
	for _, dataUnit := range s.targetData {
		fmt.Println(dataUnit.data)
	}
}


func (s *Scraper) GetUnparsedHTML() []byte {
	// for thread safety ( because we're having concurrent access)
	s.mu.RLock()
	defer s.mu.RUnlock()

	// return a copy ( also for safety because of concurrent access)
	// might remove if we ever need to change it but could just opt for a setter in that case
	cp := make([]byte, len(s.unparsedHTML))
	copy(cp, s.unparsedHTML)

	return cp
}

func (s *Scraper) Scrape() {
	// for later
}
