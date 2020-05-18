/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"
	"ubiwhere/model"

	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read N rows of data samples",
	Long: `Read: <ubiwhere read 3> to read the last three data samples. Maximum: 10.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please insert how many metrics you want to read")
			return
		}
		printLastN(args)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printLastN(args []string) {

	// Check if passed arg is a number
	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("invalid number! (Use: ubiwhere read 3)")
		return
	}

	// Limit to 10 rows
	if i > 10 {
		fmt.Println("Maximum 10!")
		return
	}

	var metrics []model.SimuData
	db := OpenDatabase()
	db.Order("id desc").Limit(args[0]).Find(&metrics)

	for _, s := range metrics {
		fmt.Printf("V1: %d | V2: %d | V3: %d | V4: %d\n", s.V1, s.V2, s.V3, s.V4)
	}
}