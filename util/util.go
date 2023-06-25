package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"strings"
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

func SetKey(key string) {
	user := os.Getenv("USER")
	sec := exec.Command("security", "add-generic-password", "-s", "pw key", "-a", user, "-w", key)
	if err := sec.Run(); err != nil {
		panic(err)
	}
}

func GetKey() (string, error) {
	user := os.Getenv("USER")
	key, err := exec.Command("security", "find-generic-password", "-w", "-s", "pw key", "-a", user).Output()
	if err != nil {
		if err.Error() == "exit status 44" {
			return "", errors.New("key not found")
		}

		panic(err)
	}
	return strings.ReplaceAll(string(key), "\n", ""), nil
}
