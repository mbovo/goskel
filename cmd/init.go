// Copyright © 2019 Manuel Bovo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/mbovo/goskel/core"
	"github.com/spf13/cobra"
)

var targetPath string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new go project",
	Long:  `A longer description here`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.ParseTemplates(targetPath); err != nil {
			fmt.Printf("%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&targetPath, "target", "t", "", "Target project directory to create")
	initCmd.MarkFlagRequired("target")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
