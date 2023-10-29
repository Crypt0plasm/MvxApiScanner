package MvxApiScanner

import (
	p "Firefly-APD"
	mt "SuperMath"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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

func AtomicUnitsToPointDecimal(Input *p.Decimal) *p.Decimal {
	DecimalDivided := mt.TruncateCustom(mt.DIVxc(Input, p.NFS("1000000000000000000")), 18)
	return DecimalDivided
}

func ConvertStringToInt64(Input string) int64 {
	Output, _ := strconv.ParseInt(Input, 10, 64)

	//if err != nil {
	//	fmt.Println("Error:", err)
	//} else {
	//	fmt.Println("Converted int64:", Output)
	//}
	return Output
}
