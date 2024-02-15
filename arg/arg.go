package arg

import (
	"fmt"
	"os"
	"path/filepath"
)

func PathToFiles(file string) string {
	dirpath, _ := os.Getwd()
	dirpath = filepath.Join(dirpath, file)
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		err := os.MkdirAll(dirpath, 0755)
		if err != nil {
			fmt.Println(err)
		}
	}
	switch file {
	case "write":
		dirpath = filepath.Join(dirpath, "data.csv")
	case "zip":
		// now := time.Now()
		// formattedTime := now.Format("2006-01-02_15:04:05")
		// archiveName := "archive_" + formattedTime + ".gz"
		dirpath = filepath.Join(dirpath, "data.csv.gz")
	}
	return dirpath
}
