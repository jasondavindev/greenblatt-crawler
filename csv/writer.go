package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/fatih/structs"
	"github.com/jasondavindev/statusinvest-crawler/config"
	"github.com/jasondavindev/statusinvest-crawler/requester_types"
)

// Write Escreve arquivo csv
func Write(d requester_types.StatusInvestResponse, outputFields config.OutputFields) {
	fmt.Printf("Salvando %d resultados em arquivo...\n", len(d))

	csvFileName := "result.csv"

	file, err := os.Create(csvFileName)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(outputFields)
	checkError("Cannot write to file", err)

	for _, value := range d {
		v := toStringSlice(outputFields, structs.Map(value))
		err := writer.Write(v)
		checkError("Cannot write to file", err)
	}

	fmt.Println("Prontinho! Seus dados estão no arquivo", csvFileName)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func toStringSlice(outputFields config.OutputFields, record map[string]interface{}) []string {
	r := make([]string, len(record))

	for i, field := range outputFields {
		r[i] = getValueAsString(record[field])
	}

	return r
}

func getValueAsString(i interface{}) string {
	switch v := i.(type) {
	case float64:
		return fmt.Sprintf("%.3f", v)
	case float32:
		return fmt.Sprintf("%.3f", v)
	case int:
		return fmt.Sprintf("%d", v)
	default:
		return fmt.Sprintf("%s", v)
	}
}
