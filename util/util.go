package util

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
)

var jsonPath = os.Getenv("HOME") + "/.data/pw/"
var jsonName = "data.json"
var jsonLocation = jsonPath + jsonName

func CopyToClipboard(pwd string) {
	pwdBuffer := bytes.Buffer{}
	pwdBuffer.Write([]byte(pwd))

	pbCopy := exec.Command("pbcopy")
	pbCopy.Stdin = &pwdBuffer

	if err := pbCopy.Start(); err != nil {
		panic(err)
	}
}

func WriteJson(passwordMap map[string]string) error {
	json, err := json.Marshal(passwordMap)
	if err != nil {
		return err
	}
	err = os.WriteFile(jsonLocation, json, 0700)
	if err != nil {
		return err
	}

	return nil
}

func ReadDataFile() ([]byte, error) {
	return os.ReadFile(jsonLocation)
}

func CreateDataFile() {
	os.MkdirAll(jsonPath, 0700)
	os.WriteFile(jsonLocation, []byte("{}"), 0700)
}

func DataFileExists() bool {
	_, err := os.Stat(jsonLocation)
	if err != nil {
		return false
	} else {
		return true
	}
}
