/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"lux-cli/utils"

	"github.com/spf13/cobra"
)

// pswCmd represents the psw command
var pswCmd = &cobra.Command{
	Use:   "psw",
	Short: "Generate randon passwords",
	Long: `Generate yours passwords:

	lux psw --t 10
	lux psw --t len_your_password`,
	Run: func(cmd *cobra.Command, args []string) {
		len, _ := cmd.Flags().GetInt("t")
		fmt.Println(utils.GeneratePassword(len))
	},
}

func init() {
	rootCmd.AddCommand(pswCmd)
	pswCmd.PersistentFlags().Int("t", 12, "Len of password")
}
