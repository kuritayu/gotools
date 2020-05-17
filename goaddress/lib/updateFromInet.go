package goaddress

import (
	"archive/zip"
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// Update 住所データを更新します。
func Update() error {
	log.Println("start download")
	if err := download(); err != nil {
		return err
	}
	log.Println("end download")

	log.Println("start unzip")
	if err := unzip(); err != nil {
		return err
	}
	log.Println("end unzip")

	log.Println("start convert")
	if err := convert(); err != nil {
		return err
	}
	log.Println("end convert")

	log.Println("start delete")
	if err := delete(); err != nil {
		return err
	}
	log.Println("end delete")
	return nil
}

func download() error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(zipf, os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	if err != nil {
		return err
	}
	if _, err = f.Write(body); err != nil {
		return err
	}
	return nil
}

func unzip() error {
	r, err := zip.OpenReader(zipf)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		buf := make([]byte, f.UncompressedSize)
		_, err = io.ReadFull(rc, buf)
		if err != nil {
			return err
		}
		if err = ioutil.WriteFile(f.Name, buf, f.Mode()); err != nil {
			return err
		}
	}
	return nil
}

func convert() error {
	sjis, err := os.Open(csvf)
	if err != nil {
		return err
	}
	defer sjis.Close()
	r := transform.NewReader(sjis, japanese.ShiftJIS.NewDecoder())
	utf8, err := os.Create(convf)
	if err != nil {
		return err
	}
	defer utf8.Close()

	tee := io.TeeReader(r, utf8)
	s := bufio.NewScanner(tee)
	for s.Scan() {
	}
	if err := s.Err(); err != nil {
		return err
	}
	return nil
}

func delete() error {
	if err := os.Remove(zipf); err != nil {
		return err
	}
	if err := os.Remove(csvf); err != nil {
		return err
	}
	return nil
}
