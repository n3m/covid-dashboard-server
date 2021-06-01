package models

import (
	"encoding/json"
	"fmt"

	"github.com/n3m/covid-dashboard-server/helpers"
)

type Case struct {
	ID                    interface{} `json:"id_registro"`
	Origen                string      `json:"origen"`
	Sector                string      `json:"sector"`
	EntidadUM             string      `json:"entidad_um"`
	Sexo                  string      `json:"sexo"`
	Entidad_Nacimiento    string      `json:"entidad_nac"`
	Entidad_Residencia    string      `json:"entidad_res"`
	Municipio_Residencia  string      `json:"municipio_res"`
	Tipo_Paciente         string      `json:"tipo_paciente"`
	Fecha_Ingreso         float64     `json:"fecha_ingreso"`
	Fecha_Sintomas        float64     `json:"fecha_sintomas"`
	Fecha_Def             *float64    `json:"fecha_def"`
	Intubado              string      `json:"intubado"`
	Neumonia              string      `json:"neumonia"`
	Edad                  float64     `json:"edad"`
	Nacionalidad          string      `json:"nacionalidad"`
	Embarazo              string      `json:"embarazo"`
	Habla_Lengua_Indig    string      `json:"habla_lengua_indig"`
	Indigena              string      `json:"indigena"`
	Diabetes              string      `json:"diabetes"`
	Epoc                  string      `json:"epoc"`
	Asma                  string      `json:"asma"`
	Inmusupr              string      `json:"inmusupr"`
	Hipertension          string      `json:"hipertension"`
	Otra_Com              string      `json:"otra_com"`
	Cardiovascular        string      `json:"cardiovascular"`
	Obesidad              string      `json:"obesidad"`
	Renal_Cronica         string      `json:"renal_cronica"`
	Tabaquismo            string      `json:"tabaquismo"`
	Otro_Caso             string      `json:"otro_caso"`
	Toma_Muestra_Lab      string      `json:"toma_muestra_lab"`
	Resultado_Lab         string      `json:"resultado_lab"`
	Toma_Muestra_Antigeno string      `json:"toma_muestra_antigeno"`
	Resultado_Antigeno    string      `json:"resultado_antigeno"`
	Clasificacion_Final   string      `json:"clasificacion_final"`
	Migrante              string      `json:"migrante"`
	Pais_Nacionalidad     string      `json:"pais_nacionalidad"`
	Pais_Origen           float64     `json:"pais_origen"`
	Uci                   string      `json:"uci"`
	Fecha_Ingreso_Yr      float64     `json:"fecha_ingreso_yr"`
	Fecha_Ingreso_Mt      float64     `json:"fecha_ingreso_mt"`
	Fecha_Ingreso_Dy      float64     `json:"fecha_ingreso_dy"`
	Fecha_Ingreso_Wk      float64     `json:"fecha_ingreso_wk"`
	Fecha_Sintomas_Yr     float64     `json:"fecha_sintomas_yr"`
	Fecha_Sintomas_Mt     float64     `json:"fecha_sintomas_mt"`
	Fecha_Sintomas_Dy     float64     `json:"fecha_sintomas_dy"`
	Fecha_Sintomas_Wk     float64     `json:"fecha_sintomas_wk"`
	Fecha_Def_Yr          *float64    `json:"fecha_def_yr"`
	Fecha_Def_Mt          *float64    `json:"fecha_def_mt"`
	Fecha_Def_Dy          *float64    `json:"fecha_def_dy"`
	Fecha_Def_Wk          *float64    `json:"fecha_def_wk"`
}

// Edit ...
func (doc *Case) Edit(value map[string]interface{}) error {
	pathError := "[models/Case.go] [Case Edit()]: %v"
	destinationMap, err := helpers.GetJSON(doc)
	if err != nil {
		return fmt.Errorf(pathError, err)
	}

	mapResult := helpers.MergeMap(destinationMap, value)
	encoded, err := json.Marshal(mapResult)
	if err != nil {
		return fmt.Errorf(pathError, err)
	}

	if err := json.Unmarshal(encoded, doc); err != nil {
		return fmt.Errorf(pathError, err)
	}

	return nil
}
