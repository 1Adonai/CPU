package main

import (
	"monitoring/metrick"
	"monitoring/write"
	"monitoring/zip"
	"time"
)

func main() {
	for {
		dataCPU := metrick.Writemetrick()
		num_lines := write.CountCsvLines()
		write.WriteinCSV(dataCPU, num_lines)
		time.Sleep(time.Second * 5)
		if num_lines >= 10000 {
			zip.CreatZIP()
		}
	}
}
