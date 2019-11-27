package main

import (
	"bufio"
	"client"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//TODO Separate Start Process
	request := flag.String("request", "request.json", "<file_name>.json")
	value := flag.String("value", "value.csv", "<file_name>.csv")
	flag.Parse()

	documents := readJson(*request)
	rows := readCsv(*value)
	var e = client.Client{}
	e.SetConnection()
	e.CreateThreads(rows)

	for _, document := range documents {
		message, _ := json.Marshal(document)
		e.MakeRequest(string(message))
	}
}

func readJson(file_name string) map[string]interface{} {
	jsonFile, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	var result map[string]interface{}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &result)
	return result
}

func readCsv(file_name string) [][]string {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))
	csv, _ := reader.ReadAll()

	return csv
}
