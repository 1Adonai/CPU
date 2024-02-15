package write

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func WriteinCSV(data []string, num_lines int) {
	var file *os.File
	var err error

	if num_lines < 100 {
		file, err = os.OpenFile("data.csv", os.O_APPEND|os.O_WRONLY, 0644)
	} else {
		file, err = os.OpenFile("data.csv", os.O_TRUNC|os.O_WRONLY, 0644)
	}

	if err != nil {
		file, err := os.Create("data.csv")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer.Flush()
}

func CountCsvLines() int {
	// Открыть файл
	data, err := ioutil.ReadFile("data.csv")
	if err != nil {
		fmt.Println(err)
	}
	// Разделить текст по символам новой строки
	lines := strings.Split(string(data), "\n")
	return len(lines)

}
