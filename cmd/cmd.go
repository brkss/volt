package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "filecrypt",
	Short: "A simple file encryption/decryption tool",
	Long:  `filecrypt is a simple CLI tool for encrypting and decrypting files using a passkey.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is provided or an invalid one is provided
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var encryptCmd = &cobra.Command{
	Use:   "e",
	Short: "Encrypt a file",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]
		passkey := args[1]
		outputFile := args[2]

		fmt.Printf("Encrypting file %s with passkey %s to %s\n", inputPath, passkey, outputFile)
		// Add your encryption logic here
	},
}

var decryptCmd = &cobra.Command{
	Use:   "d",
	Short: "Decrypt a file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]
		passkey := args[1]

		fmt.Printf("Decrypting file %s with passkey %s\n", inputPath, passkey)
		// Add your decryption logic here
	},
}
