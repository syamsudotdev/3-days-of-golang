package util

import (
	"encoding/csv"
	"os"
	"strconv"
)

func IntToString(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func WriteToCsv(data [][]string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(data)
	if err != nil {
		return err
	}

	return nil
}
