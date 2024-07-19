/*
Copyright © 2024 Allan Davis <cajun.code@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/cajun-code/semver_mgr/utility"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("commiting version file ...")
		pwd, err := os.Getwd()
		utility.Check(err)
		fileName, err := cmd.Flags().GetString("file")
		fmt.Println(fileName)
		utility.Check(err)
		r, err := git.PlainOpen(pwd)
		utility.Check(err)
		w, err := r.Worktree()
		utility.Check(err)
		_, err = w.Add(fileName)
		utility.Check(err)
		commit, err := w.Commit("Updating version number", &git.CommitOptions{
			Author: nil,
		})
		utility.Check(err)
		obj, err := r.CommitObject(commit)
		utility.Check(err)
		err = r.Push(&git.PushOptions{})
		utility.Check(err)
		fmt.Printf("Committed as %s\n", obj.Hash.String())
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	//commitCmd.Flags().StringP("file", "f", "version.json", "file to increment")
}
