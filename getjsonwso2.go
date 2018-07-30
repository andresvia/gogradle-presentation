package gogradlepresentation

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var (
	// ErrorInURL the URL was malformed
	ErrorInURL = errors.New("In URL")
	// ErrorInRequest the request failed
	ErrorInRequest = errors.New("In request")
	// ErrorInDecode response not decoded into a map[string]interface{}
	ErrorInDecode = errors.New("In decode")
)

// GetJSONWSO2 get JSON as `response` from `URL`
func GetJSONWSO2(URL string) (response map[string]interface{}, err error) {
	var req *http.Request
	if req, err = http.NewRequest("GET", URL, nil); err != nil {
		log.Print(err)
		return nil, ErrorInURL
	}
	req.Header.Set("Accept", "application/json")
	var res *http.Response
	if res, err = http.DefaultClient.Do(req); err != nil {
		log.Print(err)
		return nil, ErrorInRequest
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Print(err)
		return nil, ErrorInDecode
	}
	return response, nil
}
