/*
Copyright © 2024 Allan Davis <cajun.code@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"os"

	"github.com/cajun-code/semver_mgr/data"
	"github.com/cajun-code/semver_mgr/utility"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a new version file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		utility.Check(err)
		fmt.Println("Creating version file...")
		version := data.NewVersion()
		jsonFile, err := json.Marshal(version)
		utility.Check(err)
		err = os.WriteFile(file, jsonFile, os.ModePerm)
		utility.Check(err)
		fmt.Println("Created version file")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	//initCmd.Flags().StringP("file", "f", "version.json", "file to increment")
}
