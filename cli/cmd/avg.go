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
	"strings"

	"github.com/spf13/cobra"
)

// avgCmd represents the avg command
var avgCmd = &cobra.Command{
	Use:   "avg",
	Short: "Read AVG of one or more variables",
	Long: `Read avg: <ubiwhere read avg v1 v2> to read the AVG of the variables v1 and v2.
Vars: v1, v2, v3 and v4.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please insert the variables you want to read. (v1, v2, v3, v4)")
			return
		}
		printAvgVars(args)
	},
}

func init() {
	readCmd.AddCommand(avgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// avgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// avgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
function: Prints the average of the variables passed in args
params: args
 */
func printAvgVars(args []string) {

	// Get what variables to read
	var variables []string
	for i := 0; i < len(args); i++ {
		if strings.ToLower(args[i]) != "v1" && strings.ToLower(args[i]) != "v2" && strings.ToLower(args[i]) != "v3" && strings.ToLower(args[i]) != "v4" {
			fmt.Println("Invalid variable! Use: v1, v2, v3 and v4")
			return
		}

		// append if not exists
		_, found := Find(variables, strings.ToLower(args[i]))
		if !found {
			variables = append(variables, "AVG("+strings.ToLower(args[i])+")")
		}

	}

	// Get DB Connection
	db := OpenDatabase()
	defer db.Close()
	rows, err := db.Table("simu_data").Select(variables).Rows()
	if err != nil {
		fmt.Println("DB Error: ", err.Error())
		return
	}

	// Number of columns
	columns, _ := rows.Columns()
	colNum := len(columns)


	// Prepare a map array with the values
	var results []map[string]interface{}
	for rows.Next() {
		// Prepare to read row using Scan
		r := make([]interface{}, colNum)
		for i := range r {
			r[i] = &r[i]
		}

		// Read rows using Scan
		err = rows.Scan(r...)

		// Create a row map to store row's data
		var row = map[string]interface{}{}
		for i := range r {
			row[columns[i]] = r[i]
		}

		// Append to the final results slice
		results = append(results, row)
	}

	// Print the values
	for i, _ := range results {
		for k, v := range results[i] {
			fmt.Printf("%s: %f | ", strings.ToUpper(k), v)
		}
		fmt.Println()
	}
}