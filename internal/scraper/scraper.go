package scraper

import (
	"sync"

	"golang.org/x/net/html"
)

type Scraper struct {
	// will regroup everything needed for scraping
	// from result of url download to parsing result ect ...
	mu           sync.RWMutex
	unparsedHTML []byte

	// this might stay unused for a little while but will store the root of page
	// will also be copied and clean up into another node tree
	htmlRoot *html.Node
}

// thinking about making timeout tied to the scraper instance
func NewScraper() *Scraper {
	return &Scraper{}
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

func (s *Scraper) SetUnparsedHTML(uh []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.unparsedHTML = uh 
}

func (s *Scraper) Scrape() {
	// for later
}
