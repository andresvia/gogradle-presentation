package gogradlepresentation

import (
	"encoding/json"
	"github.com/astaxie/beego/config"
	"github.com/udistrital/administrativa_mid_api/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestObtenerInfoCoordinador(t *testing.T) {
	mockedResponses := [][]byte{}
	response := []byte{}
	bytes, _ := json.Marshal(models.ObjetoProyectoCurricular{})
	mockedResponses = append(mockedResponses, bytes)
	bytes, _ = json.Marshal(models.InformacionCoordinador{})
	mockedResponses = append(mockedResponses, bytes)
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, mockedResponses = mockedResponses[0], mockedResponses[1:]
		w.Write(response)
	}))
	defer testServer.Close()
	config := config.NewFakeConfig()
	testServerURL, _ := url.Parse(testServer.URL)
	config.Set(beegoSettingURLCRUDWSO2, testServerURL.Host)
	_, err := ObtenerInfoCoordinador("", config)
	if err != nil {
		t.Fatal(err)
	}
}

func TestErrorObtenerInformacionCoordinador(t *testing.T) {
	config := config.NewFakeConfig()
	config.Set(beegoSettingURLCRUDWSO2, " not a host ")
	_, err := ObtenerInfoCoordinador("", config)
	if err != ErrorObtenerInformacionCoordinador {
		t.Fail()
	}
}

func TestErrorErrorModeloCoordinador(t *testing.T) {
	mockedResponses := [][]byte{}
	response := []byte{}
	bytes, _ := json.Marshal(models.ObjetoProyectoCurricular{})
	mockedResponses = append(mockedResponses, bytes)
	bytes, _ = json.Marshal(map[string]string{"CarreraSniesCollection": "bad map"})
	mockedResponses = append(mockedResponses, bytes)
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, mockedResponses = mockedResponses[0], mockedResponses[1:]
		w.Write(response)
	}))
	defer testServer.Close()
	config := config.NewFakeConfig()
	testServerURL, _ := url.Parse(testServer.URL)
	config.Set(beegoSettingURLCRUDWSO2, testServerURL.Host)
	_, err := ObtenerInfoCoordinador("", config)
	if err != ErrorModeloCoordinador {
		t.Fail()
	}
}
