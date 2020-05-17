package goaddress

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// CodeToAdd 郵便番号を住所に変換します。
func CodeToAdd(code string) error {
	replaced := strings.Replace(code, "-", "", -1)
	reader, err := readData()
	if err != nil {
		return err
	}

	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		if line[2] == replaced {
			fmt.Printf("%v %v%v%v\n", line[2], line[6], line[7], line[8])
		}
	}
	return nil
}

// AddToCode 住所を郵便番号に変換します。
func AddToCode(add string) error {
	reader, err := readData()
	if err != nil {
		return err
	}

	r := regexp.MustCompile(add)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		target := fmt.Sprintf("%v%v%v", line[6], line[7], line[8])
		if r.MatchString(target) {
			fmt.Printf("%v %v%v%v\n", line[2], line[6], line[7], line[8])
		}
	}
	return nil
}

func readData() (*csv.Reader, error) {
	f, err := os.Open(convf)
	if err != nil {
		return nil, err
	}

	return csv.NewReader(f), nil
}
