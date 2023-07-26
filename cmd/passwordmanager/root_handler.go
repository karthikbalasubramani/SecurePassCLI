package main

import (
	"fmt"
	"encoding/json"
	"github.com/bkarthik338/password-manager/pkg/common"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"github.com/spf13/cobra"
)

var rootFileName string = ".rootuser.json"

func SigninHandler(cmd *cobra.Command, args []string){
	fmt.Println("The login is required to access Password Manager:")
	username := Prompt("Root Username: ")
	password := Prompt("Root Password: ")
	fileData, err := ioutil.ReadFile(".rootuser.json")
	var storedCredentials common.RootUserEntry
	err = json.Unmarshal(fileData, &storedCredentials)
	if err != nil {
		log.Fatal("Failed to unmarshal user credentials:", err)
	}
	if username != storedCredentials.Username {
		fmt.Println("Invalid Username ",username)
		SigninHandler(cmd, args)
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedCredentials.Password), []byte(password))
	if err != nil {
		fmt.Println("Invalid Password")
		SigninHandler(cmd, args)
	}
	fmt.Println("Login Successful: ",username)
}


func SignupHandler(cmd *cobra.Command, args []string){
	username := Prompt("Username: ")
	password := Prompt("Password: ")
	//Creating new password instance
	rootUserEntry := common.RootUserEntry{
		Username: username,
		Password: password,
	}
	//Validation for non-null values for username and password
	if rootUserEntry.Username == "" || rootUserEntry.Password == "" {
		fmt.Println("Username and password are mandatory fields")
		// Call the handler recursively to prompt for details again
		SignupHandler(cmd, args)
		return
	}
	rootUserEntry.Password, _ = common.HashPassword(rootUserEntry.Password)
	// Marshal the updated entries to JSON
	entriesJSON, err := json.Marshal(rootUserEntry)
	if err != nil {
		log.Fatal("Failed to marshal entries to JSON:", err)
	}
	// Write the updated entries to the file
	err = ioutil.WriteFile(rootFileName, entriesJSON, 0644)
	if err != nil {
		log.Fatal("Unable to write entries to file:", err)
	}
	fmt.Println("Successfully added Root user")
}