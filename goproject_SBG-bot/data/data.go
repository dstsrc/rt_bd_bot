package data

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	s "strings"
)

type Reader struct {
	filePath string
}

const (
//	n_nomer = 3

)

func NewReader(file string) Reader {
	return Reader{filePath: file}
}

func ReadFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка чтения")
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Ошибка при чтении записи в переменную")
		return nil, err
	}

	return records, nil
}

func (r Reader) ReadFile_2() ([][]string, error) {
	f, err := ioutil.ReadFile(r.filePath)
	if err != nil {
		fmt.Println("ошибка")
	}

	t := s.Split(string(f), "\n")

	z := make([]string, len(t))
	for i, v := range t {
		for _, w := range v {
			if w == 13 {
				break
			} else {
				z[i] = z[i] + string(w)
				//	fmt.Println(w, string(w))
			}
		}
	}
	records := make([][]string, 0, 10)

	for i, _ := range z {

		if len(z[i]) != 0 {
			records = append(records, s.Split(z[i], ","))
			//	fmt.Println(r[i], " ", len(r[i]))
		}
	}

	return records, nil
}

func WriteFile(filePath string, records [][]string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	csvWriter := csv.NewWriter(f)
	err = csvWriter.WriteAll(records)
	if err != nil {
		return err
	}

	return nil
}

func WriteFile_2(filePath string, records [][]string) error {
	var text string
	for i, _ := range records {
		for j, _ := range records[i] {
			x := ","
			if len(records[i])-1 == j {
				x = ""
			}
			text = text + records[i][j] + x
		}
		y := "\n"
		if len(records)-1 == i {
			y = ""
		}
		text = text + y
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		return err
	}
	defer f.Close()
	f.WriteString(text)

	return nil
}
