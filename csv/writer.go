package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/fatih/structs"
	"github.com/jasondavindev/statusinvest-crawler/requester"
)

// Write Escreve arquivo csv
func Write(d requester.StatusInvestResponse) {
	fmt.Println("Salvando dados em arquivo...")

	csvFileName := "result.csv"

	file, err := os.Create(csvFileName)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(structs.Names(&requester.StatusInvestResponseItem{}))
	checkError("Cannot write to file", err)

	for _, value := range d {
		v := parseToStringSlice(structs.Values(value))
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

func parseToStringSlice(v []interface{}) []string {
	r := make([]string, len(v))

	for i, v := range v {
		r[i] = getValueAsString(v)
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
