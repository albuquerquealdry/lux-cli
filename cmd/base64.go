/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"lux-cli/utils"

	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Base64 Encode and Decode",
	Long: `Exemple Encode and decode strings:
	encode: lux base64 --e "you string"
	decode: lux base64 --d "you string"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		encodeStr, _ := cmd.Flags().GetString("e")
		decodeStr, _ := cmd.Flags().GetString("d")
		if encodeStr != "" {
			encode := utils.EncodeString(encodeStr)
			cmd.Println(encode)
		} else if decodeStr != "" {
			decode := utils.DecodeString(decodeStr)
			cmd.Println(decode)
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.PersistentFlags().String("e", "", "Encode Base64")
	base64Cmd.PersistentFlags().String("d", "", "Decode Base64")

}
