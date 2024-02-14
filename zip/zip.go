package zip

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"time"
)

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
