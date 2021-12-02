package supporting

import (
	"bufio"
	"errors"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
)

// логированние программы
func CreateLogger(logger *log.Logger) (*log.Logger, *os.File) {
	logfile, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	l, err := os.OpenFile(logfile+"/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger = log.New(l, "customLog ", log.LstdFlags)
	logger.SetFlags(log.LstdFlags)
	return logger, l
}

func CheckSum(errC chan error, CheckSum string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	m, err := os.Open(usr.HomeDir + "/go/src/magma/magma.go")

	r := bufio.NewReader(m)
	b, err := ioutil.ReadAll(r)

	table := crc32.MakeTable(0xD5828281)

	sum := crc32.Checksum(b, table)
	s := strconv.Itoa(int(sum))
	if s != CheckSum {
		err = errors.New("binary file is not legitimate")
	}
	errC <- err
}
func BegincheckSum() (string, error){
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	m, err := os.Open(usr.HomeDir + "/go/src/magma/magma.go")

	r := bufio.NewReader(m)
	b, err := ioutil.ReadAll(r)

	table := crc32.MakeTable(0xD5828281)

	sum := crc32.Checksum(b, table)
	s := strconv.Itoa(int(sum))
	return s, err
}
