package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	for {
		dataCPU := writemetrick()
		num_lines := countCsvLines()
		if num_lines >= 10000 {
			creatZIP()
		}
		writeinCSV(dataCPU, num_lines)
		time.Sleep(time.Second * 5)

	}
}

// Получение метрик
func getCPUSample() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}

// Вычисление и запись в переменную
func writemetrick() []string {
	idle0, total0 := getCPUSample()
	time.Sleep(3 * time.Second)
	idle1, total1 := getCPUSample()

	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks
	formatCPU := math.Trunc(cpuUsage*100) / 100
	dataCPU := []string{
		strconv.FormatFloat(formatCPU, 'f', 2, 64),
		strconv.FormatFloat(totalTicks-idleTicks, 'f', 2, 64),
		strconv.FormatFloat(totalTicks, 'f', 2, 64),
	}
	return dataCPU
}

// Запись в CSV
func writeinCSV(data []string, num_lines int) {
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

		fmt.Println("Файл data.csv создан")
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

func countCsvLines() int {
	// Открыть файл
	data, err := ioutil.ReadFile("data.csv")
	if err != nil {
		fmt.Println(err)
	}

	// Разделить текст по символам новой строки
	lines := strings.Split(string(data), "\n")
	return len(lines)

}

func creatZIP() {
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
