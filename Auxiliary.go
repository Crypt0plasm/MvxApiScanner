package MvxApiScanner

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// OnPage Function Snapshots Link and returns String
func OnPage(Link string) string {
	res, err := http.Get(Link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

// MyCopy Function copies a  file from a source to a destination
func MyCopy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
