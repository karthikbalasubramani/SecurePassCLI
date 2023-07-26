package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

//In this rootCmd variable all the new commands are added.
var rootCmd = &cobra.Command{
	Use:   "passwordmanager",
	Short: "A password manager CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Password Manager CLI!")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
	StartShell()
}