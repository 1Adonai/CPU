package zip

import (
	"compress/gzip"
	"fmt"
	"io"
	"monitoring/arg"
	"os"
)

func CreatZIP() {
	pathzip := arg.PathToFiles("zip")
	gzFile, err := os.Create(pathzip)
	if err != nil {
		fmt.Println(err)
		return
	}
	gzipWriter := gzip.NewWriter(gzFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	pathdata := arg.PathToFiles("write")
	csvFile, err := os.Open(pathdata)
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
