package gogradlepresentation

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/udistrital/administrativa_mid_api/models"
	"log"
)

var (
	// ErrorObtenerInformacionProyectoCurricular al obtener proyecto curricular
	ErrorObtenerInformacionProyectoCurricular = errors.New("Información de proyecto curricular")
	// ErrorModeloProyectoCurricular al mapear respuesta a modelo
	ErrorModeloProyectoCurricular = errors.New("Modelo de proyecto curricular")
)

// ObtenerInformacionCoordinador en dos pasos, primero dependencia, luego carrera
func ObtenerInformacionCoordinador(urlIDDependenciaOikos, urlCarreraSniesFmt string) (response map[string]interface{}, err error) {
	if response, err = GetJSONWSO2(urlIDDependenciaOikos); err != nil {
		log.Print(err)
		return nil, ErrorObtenerInformacionProyectoCurricular
	}
	var proyectoCurricular models.ObjetoProyectoCurricular
	err = mapstructure.Decode(response, &proyectoCurricular)
	if err != nil {
		log.Print(err)
		return nil, ErrorModeloProyectoCurricular
	}
	urlIDProyectoSnies := fmt.Sprintf(urlCarreraSniesFmt, proyectoCurricular.Homologacion.IDSnies)
	return GetJSONWSO2(urlIDProyectoSnies)
}
