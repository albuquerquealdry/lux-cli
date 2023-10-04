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
	lux psw --t len_your_password
	lux psw --l 5
	lux psw --t len logic password`,

	Run: func(cmd *cobra.Command, args []string) {
		len, _ := cmd.Flags().GetInt("t")
		logic, _ := cmd.Flags().GetInt("l")
		if len > 8 {
			fmt.Println(utils.GeneratePassword(len))
		} else {
			fmt.Println(utils.GenerateLogicPassword(logic))
		}

	},
}

func init() {
	rootCmd.AddCommand(pswCmd)
	pswCmd.PersistentFlags().Int("t", 2, "Len of password")
	pswCmd.PersistentFlags().Int("l", 1, "Len of password")
}
