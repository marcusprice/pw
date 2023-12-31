package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/marcusprice/pw/encryption"
	"github.com/marcusprice/pw/store"
	"github.com/marcusprice/pw/util"
)

const UNKNOWN_COMMAND_MESSAGE = "unknown action, list of available actions"

var passwordStore *store.Store

func main() {
	key, err := util.GetKey()
	if err != nil {
		key = encryption.GenerateKey()
		util.SetKey(key)
	}
	passwordData := getPasswordData()
	passwordStore = store.NewPasswordStore(passwordData)

	var action string = ""
	argsPresent := len(os.Args) - 1
	if argsPresent < 1 {
		printHelpMenu()
		return
	} else {
		action = os.Args[1]
	}

	if argsPresent == 1 {

		if action == "help" {
			printHelpMenu()
		} else {
			printUnknownCommand()
		}

	} else if argsPresent == 2 {

		service := encryption.Encrypt(os.Args[2], key)
		if action == "get" {
			getPassword(service, key)
		} else if action == "delete" {
			deletePassword(service)
		} else {
			printUnknownCommand()
		}

	} else if argsPresent == 3 {

		service := os.Args[2]
		encryptedService := encryption.Encrypt(service, key)
		encryptedPassword := encryption.Encrypt(os.Args[3], key)

		if action == "new" {
			_, err := newPassword(encryptedService, encryptedPassword)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("new password saved for" + service)
		} else if action == "edit" {
			editPassword(encryptedService, encryptedPassword)
			fmt.Println(service + " password updated")
		} else {
			printUnknownCommand()
		}
	}
}

func newPassword(service string, pwd string) (ok bool, err error) {
	if passwordStore.ServiceExists(service) {
		return false, errors.New("service already exists")
	}
	passwordStore.Write(service, pwd)
	util.WriteJson(passwordStore.GetStore())
	return true, nil
}

func getPassword(service string, key string) {
	if pwd, err := passwordStore.Read(service); err != nil {
		fmt.Println(err)
	} else {
		decryptedPassword := encryption.Decrypt(pwd, key)
		util.CopyToClipboard(decryptedPassword)
		fmt.Println("password copied to clipboard!")
	}
}

func editPassword(service string, pwd string) {
	passwordStore.Write(service, pwd)
	util.WriteJson(passwordStore.GetStore())
}

func deletePassword(service string) {
	passwordStore.Delete(service)
	util.WriteJson(passwordStore.GetStore())
}

func printHelpMenu() {
	fmt.Println("Available Commands:")
	fmt.Println("new - Save a new password")
	fmt.Println("get - Get a password")
	fmt.Println("edit - Edit an existing password")
	fmt.Println("delete - Delete a password")
}

func printUnknownCommand() {
	fmt.Println(UNKNOWN_COMMAND_MESSAGE)
	printHelpMenu()
}

func getPasswordData() store.PasswordData {
	var passwordData store.PasswordData

	if util.DataFileExists() {
		file, err := util.ReadDataFile()
		if err != nil {
			panic(err)
		} else {
			json.Unmarshal(file, &passwordData)
		}
	} else {
		util.CreateDataFile()
		passwordData = make(map[string]string)
	}
	return passwordData
}
