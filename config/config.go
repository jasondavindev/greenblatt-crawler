package config

import (
	"bytes"
	"io/ioutil"
	"log"

	"go.uber.org/config"
)

// Filters Filtros disponiveis
type Filters struct {
	Dividendo    [2]float32
	PL           [2]float32
	DivLiqEbit   [2]float32
	DivLiqPatLiq [2]float32
	Roe          [2]float32
	Roic         [2]float32
	Roa          [2]float32
	LiqMedDiaria [2]float32
}

// CfgFactory Carrega todos os filtros
func CfgFactory(configPath string) Filters {
	yamlFile, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Fatal(err)
	}

	yamlFileReader := bytes.NewReader(yamlFile)

	provider, err := config.NewYAML(config.Source(yamlFileReader))
	if err != nil {
		log.Fatal(err)
	}

	var c Filters

	if err := provider.Get("filtros").Populate(&c); err != nil {
		log.Fatal(err)
	}

	return c
}
