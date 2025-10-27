package scraper

import (
	"sync"
)

type Scraper struct {
	// to be implemented
	// will regroup everything needed for scraping
	// from result of url download to parsing result ect ...

	mu sync.RWMutex
	unparsedHTML []byte
	targetData []DataUnit
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
