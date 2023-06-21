package util

import (
	"encoding/json"
	"os"
	"os/exec"
)

var jsonPath = os.Getenv("HOME") + "/.data/pw/"
var jsonName = "data.json"
var jsonLocation = jsonPath + jsonName

func CopyToClipboard(pwd string) error {
	echo := exec.Command("echo", pwd)
	pbCopy := exec.Command("pbcopy")
	pipe, _ := echo.StdoutPipe()
	defer pipe.Close()

	pbCopy.Stdin = pipe

	echo.Start()
	pbCopy.Start()

	return nil
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
