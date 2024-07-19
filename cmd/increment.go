/*
Copyright © 2024  Allan Davis <cajun.code@gmail.com>
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

// incrementCmd represents the increment command
var incrementCmd = &cobra.Command{
	Use:   "increment",
	Short: "increse the given version field by 1",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("increment called")
		field, _ := cmd.Flags().GetString("mode")
		fileName, err := cmd.Flags().GetString("file")
		utility.Check(err)
		fmt.Printf("Opening version %s ...\n", fileName)
		jsonFile, err := os.ReadFile(fileName)
		utility.Check(err)
		var version data.Version
		err = json.Unmarshal(jsonFile, &version)
		utility.Check(err)
		fmt.Printf("Incrementing %s by 1\n", field)
		version.Increment(field)
		jsonFile, err = json.Marshal(version)
		utility.Check(err)
		fmt.Println("Updating version file...")
		err = os.WriteFile(fileName, jsonFile, os.ModePerm)
		utility.Check(err)
		fmt.Println("updated version file")
	},
}

func init() {
	rootCmd.AddCommand(incrementCmd)

	incrementCmd.Flags().StringP("mode", "m", "build", "field to increment")
	//incrementCmd.Flags().StringP("file", "f", "version.json", "file to increment")
}
