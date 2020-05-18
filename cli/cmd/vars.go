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
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

// varsCmd represents the vars command
var varsCmd = &cobra.Command{
	Use:   "vars",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 1 {
			fmt.Println("Please insert how many metrics you want to read + the variables you want")
			return
		}
		printLastNVars(args)
	},
}

func init() {
	readCmd.AddCommand(varsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// varsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// varsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printLastNVars(args []string) {

	// Check if passed first arg is a number
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

	// Get what variables to read
	var variables []string
	for i := 1; i < len(args); i++ {
		if strings.ToLower(args[i]) != "v1" && strings.ToLower(args[i]) != "v2" && strings.ToLower(args[i]) != "v3" && strings.ToLower(args[i]) != "v4" {
			fmt.Println("Invalid variable! Use: v1, v2, v3 and v4")
			return
		}
		//fmt.Println("->", strings.ToLower(args[i]))
		variables = append(variables, strings.ToLower(args[i]))
	}

	//var samples []model.SimuData
	db := OpenDatabase()
	rows, err := db.Table("simu_data").Order("id desc").Select(variables).Limit(i).Rows()

	for rows.Next() {

	}

}