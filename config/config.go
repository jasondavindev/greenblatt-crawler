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

// CfgFactory Carrega todos os filtros
func CfgFactory(configPath string) Filters {
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

	var c Filters

	if err := provider.Get("filtros").Populate(&c); err != nil {
		log.Fatal(err)
	}

	return c
}
