package gogradlepresentation

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/mitchellh/mapstructure"
	"github.com/udistrital/administrativa_mid_api/models"
	"log"
)

const (
	beegoSettingURLCRUDWSO2        = "UrlcrudWSO2"
	beegoSettingNSCRUDHomologacion = "NscrudHomologacion"
	beegoSettingNSCRUDAcademica    = "NscrudAcademica"
)

var (
	// ErrorObtenerInformacionCoordinador al obtener coordinador
	ErrorObtenerInformacionCoordinador = errors.New("Obteniendo información de coordinador")
	// ErrorModeloCoordinador al mapear respuesta a modelo
	ErrorModeloCoordinador = errors.New("Modelo de coordinador")
)

// ObtenerInfoCoordinador obtain info or fail
func ObtenerInfoCoordinador(idDependenciaOikos string, beegoConfiger config.Configer) (response models.InformacionCoordinador, err error) {
	// beegoConfiger = beego.AppConfig
	var apiResponse map[string]interface{}
	// protocol should not be hardcoded in http://
	urlIDDependenciaOikos := fmt.Sprintf("http://%s/%s/proyecto_curricular_oikos/%s", beegoConfiger.String(beegoSettingURLCRUDWSO2), beegoConfiger.String(beegoSettingNSCRUDHomologacion), idDependenciaOikos)
	urlCarreraSniesFmt := fmt.Sprintf("http://%s/%s/carrera_snies/%%s", beegoConfiger.String(beegoSettingURLCRUDWSO2), beegoConfiger.String(beegoSettingNSCRUDAcademica))
	if apiResponse, err = ObtenerInformacionCoordinador(urlIDDependenciaOikos, urlCarreraSniesFmt); err != nil {
		log.Print(err)
		return response, ErrorObtenerInformacionCoordinador
	}
	err = mapstructure.Decode(apiResponse, &response)
	if err != nil {
		log.Print(err)
		return response, ErrorModeloCoordinador
	}
	return response, nil
}
