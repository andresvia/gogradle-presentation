package gogradlepresentation

import (
	"encoding/json"
	"fmt"
	"github.com/udistrital/administrativa_mid_api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestObtenerInformacionCoordinador(t *testing.T) {
	attribute := "attribute1"
	value := "value1"
	mockedResponses := [][]byte{}
	response := []byte{}
	bytes, _ := json.Marshal(models.ObjetoProyectoCurricular{})
	mockedResponses = append(mockedResponses, bytes)
	bytes, _ = json.Marshal(map[string]string{attribute: value})
	mockedResponses = append(mockedResponses, bytes)
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, mockedResponses = mockedResponses[0], mockedResponses[1:]
		w.Write(response)
	}))
	defer testServer.Close()
	serverResponse, err := ObtenerInformacionCoordinador(testServer.URL, testServer.URL+"/%s")
	if err != nil {
		t.Fatal(err)
	}
	attributeValue := serverResponse[attribute].(string)
	if attributeValue != value {
		t.Fail()
	}
}

func TestErrorObtenerInformacionProyectoCurricular(t *testing.T) {
	_, err := ObtenerInformacionCoordinador("hxxt:// not a url /", "hxxt:// not a url /")
	if err != ErrorObtenerInformacionProyectoCurricular {
		t.Fail()
	}
}

func TestErrorModeloProyectoCurricular(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"Homologacion":"bad map"}`)
	}))
	_, err := ObtenerInformacionCoordinador(testServer.URL, testServer.URL+"/%s")
	if err != ErrorModeloProyectoCurricular {
		t.Fail()
	}
}
