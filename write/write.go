package write

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"monitoring/arg"
	"os"
	"strings"
)

func WriteinCSV(data []string, num_lines int) {
	var file *os.File
	var err error
	path := arg.PathToFiles("write")
	if num_lines < 10000 {
		file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	} else {
		file, err = os.OpenFile(path, os.O_TRUNC|os.O_WRONLY, 0644)
	}

	if err != nil {
		file, err := os.Create(path)
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
	path := arg.PathToFiles("write")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	// Разделить текст по символам новой строки
	lines := strings.Split(string(data), "\n")
	return len(lines)

}
