package write

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func WriteinCSV(data []string, num_lines int) {
	var file *os.File
	var err error

	if num_lines < 10000 {
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

func CreatZIP() {
	now := time.Now()
	formattedTime := now.Format("2006-01-02_15:04:05") // Форматирование по шаблону
	archiveName := "archive_" + formattedTime + ".gz"
	gzFile, err := os.Create(archiveName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Сжимаем файл CSV
	gzipWriter := gzip.NewWriter(gzFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.Copy(gzipWriter, csvFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	csvFile.Close()
	gzipWriter.Close()
	gzFile.Close()
}
