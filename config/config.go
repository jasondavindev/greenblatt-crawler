package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"go.uber.org/config"
)

// Filters Filtros disponiveis
type Filters map[string]interface{}

type Grouping struct {
	Class string
	Top   int
}

type ConfigSetup struct {
	Filters  Filters
	Grouping Grouping
}

// CfgFactory Carrega todos os filtros
func CfgFactory(configPath string) ConfigSetup {
	fmt.Println("Carregando arquivo contendo os filtros...")

	yamlFile, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Fatal(err)
	}

	yamlFileReader := bytes.NewReader(yamlFile)

	provider, err := config.NewYAML(config.Source(yamlFileReader))
	if err != nil {
		log.Fatal(err)
	}

	var filters Filters
	var grouping Grouping

	if err := provider.Get("filtros").Populate(&filters); err != nil {
		log.Fatal(err)
	}

	if err := provider.Get("agrupamento").Populate(&grouping); err != nil {
		log.Fatal(err)
	}

	return ConfigSetup{
		Filters:  filters,
		Grouping: grouping,
	}
}
