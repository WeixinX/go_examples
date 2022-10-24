package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func readCSV(input string) ([][]string, error) {
	f, err := os.Open(input)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	datas := make([][]string, 0)
	r := csv.NewReader(f)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return [][]string{}, err
		}

		datas = append(datas, row)
	}
	return datas, nil
}

func writeCSV(datas [][]string, output string) error {
	f, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	err = w.WriteAll(datas)
	if err != nil {
		return err
	}
	return nil
}

// 重排第一列 1, 2, ..., n-1, n
func handleCSV(datas [][]string) {
	n := len(datas)
	for i := 1; i <= n; i++ {
		datas[i-1][0] = strconv.Itoa(i)
	}
}

func main() {
	input := "./test.csv"
	output := "./test_out.csv"
	datas, err := readCSV(input)
	if err != nil {
		fmt.Println(err)
	}
	handleCSV(datas)
	err = writeCSV(datas, output)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("completed")
}
