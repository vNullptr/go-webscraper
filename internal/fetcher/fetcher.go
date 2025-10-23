package fetcher

import (
	"net/http" // http requests 
	//"net/url" // url handling and type
	"io" // to read the body
	"errors" // for error handling
)

// is there a point in having this for organization purposes??
type HttpClient struct {
	Url string
	Method string
}

// need to use http.NewRequest for other methods and custom headers 
// bonus : add timeout handling
func (h *HttpClient) Fetch() (string, error){
	
	resp, err := http.Get(h.Url)

	if err != nil {

		return "", errors.New("Failed to fetch URL: " + err.Error())
		defer resp.Body.Close()

	} else {
		if ( resp.StatusCode == http.StatusOK ) {

			body, _ := io.ReadAll(resp.Body)
			return string(body), nil

		} else {
			return "", errors.New("Error HTTP status: " + resp.Status)
		}
	}

	return "", errors.New("Unknown error")

}
