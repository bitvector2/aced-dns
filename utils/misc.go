package utils

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"

	log "github.com/golang/glog"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateFile(filename string, newData []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, newData, perm)
}

func UpdateFile(filename string, newData []byte, perm os.FileMode) (bool, error) {
	var err error
	var oldData []byte

	// read error occurred thus createFile should be called
	oldData, err = ioutil.ReadFile(filename)
	if err != nil {
		return false, err
	}

	if bytes.Compare(oldData, newData) != 0 {
		err = ioutil.WriteFile(filename, newData, perm)
		return true, err
	}

	return false, nil
}

func RunRndc() {
	rndcCmd := exec.Command("rndc", "reload")
	text, err := rndcCmd.Output()
	Check(err)
	log.Info(text)
}
