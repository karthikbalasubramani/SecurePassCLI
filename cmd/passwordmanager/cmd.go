package main

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(helpCmd)
	rootCmd.AddCommand(signupCmd)
	rootCmd.AddCommand(signincmd)
	rootCmd.AddCommand(exitCmd)
	rootCmd.AddCommand(getCmd)
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Help about any command",
	Run:   HelpHandler,
}

var exitCmd = &cobra.Command{
	Use: "exit",
	Short: "Exit from Interactive Shell",
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new password entry",
	Run:   AddHandler,
}

var signupCmd = &cobra.Command{
	Use: "signup",
	Short: "Create a new root user",
	Run: SignupHandler,
}

var signincmd = &cobra.Command{
	Use: "signin",
	Short: "Root User Login",
	Run: SigninHandler,
}

var getCmd = &cobra.Command{
	Use: "get",
	Short: "Get Password Entry",
	Run: GetHandler,
}