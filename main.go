package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/marcusprice/pw/encryption"
	"github.com/marcusprice/pw/store"
	"github.com/marcusprice/pw/util"
)

const UNKNOWN_COMMAND_MESSAGE = "unknown action, list of available actions"

var passwordStore store.Store

func main() {
	passwordStore.Init()

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

		service, _ := encryption.Encrypt(os.Args[2])
		if action == "get" {
			getPassword(service)
		} else if action == "delete" {
			deletePassword(service)
		} else {
			printUnknownCommand()
		}

	} else if argsPresent == 3 {

		service := os.Args[2]
		encryptedService, _ := encryption.Encrypt(service)
		pwd, _ := encryption.Encrypt(os.Args[3])

		if action == "new" {
			_, err := saveNewPassword(encryptedService, pwd)
			if err != nil {
				fmt.Println(err)
			}
		} else if action == "edit" {
			editPassword(encryptedService, pwd)
			fmt.Println(service + " password updated")
		} else {
			printUnknownCommand()
		}
	}
}

func saveNewPassword(service string, pwd string) (ok bool, err error) {
	if passwordStore.ServiceExists(service) {
		return false, errors.New("service already exists")
	}
	passwordStore.Add(service, pwd)
	util.WriteJson(passwordStore.GetStore())
	return true, nil
}

func getPassword(service string) {
	pwd, err := passwordStore.Get(service)
	if err != nil {
		fmt.Println(err)
		return
	}

	decryptedPassword, _ := encryption.Decrypt(pwd)

	err = util.CopyToClipboard(decryptedPassword)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("password copied to clipboard!")
}

func editPassword(service string, pwd string) {
	passwordStore.Add(service, pwd)
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
