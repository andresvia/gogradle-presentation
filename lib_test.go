package gogradlepresentation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessfulResponse(t *testing.T) {
	attribute := "attribute1"
	value := "value1"
	mockedResponse, _ := json.Marshal(map[string]string{attribute: value})
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(mockedResponse)
	}))
	defer testServer.Close()
	serverResponse, err := GetJSONWSO2(testServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	attributeValue := serverResponse[attribute].(string)
	if attributeValue != value {
		t.Fail()
	}
}

func TestErrorInURL(t *testing.T) {
	_, err := GetJSONWSO2("hxxt:// not a url /")
	if err != ErrorInURL {
		t.Fail()
	}
}

func TestErrorInRequest(t *testing.T) {
	_, err := GetJSONWSO2("https://server:port/path?query")
	if err != ErrorInRequest {
		t.Fail()
	}
}

func TestErrorInDecode(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Not JSON")
	}))
	defer testServer.Close()
	_, err := GetJSONWSO2(testServer.URL)
	if err != ErrorInDecode {
		t.Fatal(err)
	}
}
