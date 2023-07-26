package main

import (
	"fmt"
	"encoding/json"
	"github.com/bkarthik338/password-manager/pkg/common"
	"io/ioutil"
	"github.com/spf13/cobra"
	"log"
)

var Prompt = common.Prompt

var passwordEntriesFileName string = ".passwordentries.json"

func HelpHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Available commands:")
	for _, c := range rootCmd.Commands() {
		if !c.Hidden && (c.Name() != "completion" && c.Name() != "help"){
			fmt.Printf("  %-15s %s\n", c.Name(), c.Short)
		}
	}
}

func AddHandler(cmd *cobra.Command, args []string) {
	applicationname := Prompt("Name: ")
	username := Prompt("Username: ")
	password := Prompt("Password: ")
	//Creating new password instance
	passwordEntry := common.PasswordEntry{
		ApplicationName: applicationname,
		Username: username,
		Password: password,
	}
	//Validation for non-null values for username and password
	if passwordEntry.Username == "" || passwordEntry.Password == "" {
		fmt.Println("Username and password are mandatory fields")
		// Call the handler recursively to prompt for details again
		AddHandler(cmd, args)
		return
	}
	// Read existing entries from the file, if any
	existingEntries, err := common.ReadEntriesFromFile(passwordEntriesFileName)
	if err != nil {
		log.Fatal("Failed to read existing entries from file:", err)
	}
	// Append the new entry to the existing entries
	entries := append(existingEntries, passwordEntry)
	// Marshal the updated entries to JSON
	entriesJSON, err := json.Marshal(entries)
	if err != nil {
		log.Fatal("Failed to marshal entries to JSON:", err)
	}
	// Write the updated entries to the file
	err = ioutil.WriteFile(passwordEntriesFileName, entriesJSON, 0644)
	if err != nil {
		log.Fatal("Unable to write entries to file:", err)
	}
	fmt.Println("Successfully updated password entry for: ", applicationname)
}


func GetHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Please select any option to retreive password entry:")
	option := Prompt("1). Username.\n2). Application Name\n")
	switch option {
		case "1":
			username := Prompt("Enter the Username: ")
			entries, err := common.GetPasswordEntry(passwordEntriesFileName, username, "")
			if err != nil {
				fmt.Println(err)
				GetHandler(cmd, args)
			}
			entriesJson, _ := json.Marshal(entries)
			fmt.Println("%v", entriesJson)
			// for _, entry := range entriesJson{
			// 	fmt.Println("Application Name: ", entry["applicationname"])
			// 	fmt.Println("Username: ", entry["username"])
			// 	fmt.Println("Password: ", entry["password"])
			// }
		case "2":
			applicationname := Prompt("Enter the Application Name: ")
			entries, err := common.GetPasswordEntry(passwordEntriesFileName, "", applicationname)
			if err != nil {
				fmt.Println(err)
				GetHandler(cmd, args)
			}
			entriesJson, _ := json.Marshal(entries)
			// for _, entry := range entriesJson {
			// 	fmt.Println("Application Name: ", entry["applicationname"])
			// 	fmt.Println("Username: ", entry["username"])
			// 	fmt.Println("Password: ", entry["password"])	
			// }		
			fmt.Println("%v", entriesJson)
	}
}
