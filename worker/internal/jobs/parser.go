package jobs

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Parse(file string, c chan [][]string) {
	fileContent, fileErr := os.Open(file)

	if fileErr != nil {
		fmt.Println(fileErr.Error())
	}
	defer fileContent.Close()

	csvReader := csv.NewReader(fileContent)
	fileData, readErr := csvReader.ReadAll()

	if readErr != nil {
		fmt.Println(readErr.Error())
	}

	c <- fileData
}
