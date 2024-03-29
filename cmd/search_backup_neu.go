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
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/dominik-robert/emiro/emironetwork"
	"github.com/dominik-robert/emiro/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search and List the available command",
	Long: `With search you can search to all the existing commands. You can search for a specific command or list all commands. An alias for search is list

Examples:

# List all existing commands. 
emiro search -a 

# List two existings commands
emiro search -a -c 2 

# List the first 5 commands that match the "kubernetes" query
emiro search kubernetes -c 5`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		count, _ := cmd.Flags().GetInt32("count")
		emiroHost := viper.GetString("emiroHost")
		emiroPort := viper.GetInt("emiroPort")
		verbose, _ := cmd.Flags().GetBool("verbose")
		var conn *grpc.ClientConn
		var err error

		conn, err = grpc.Dial(emiroHost+":"+fmt.Sprint(emiroPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Could not connect to server: %s", err)
		}
		defer conn.Close()

		if verbose {
			log.Println("Connect to", emiroHost+":"+fmt.Sprint(emiroPort))
		}

		c := emironetwork.NewEmiroClient(conn)

		queryString := args[0]

		query := emironetwork.Query{
			Query: queryString,
			Count: count,
		}

		response, err := c.SendQuery(context.Background(), &query)

		if err != nil {
			log.Fatalf("Error when calling SendQuery: %s", err)
		}

		outputFormat, _ := cmd.Flags().GetString("output")
		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

		switch outputFormat {
		case "detailed":
			for _, value := range response.QueryAnswers {
				fmt.Fprintln(writer, "ID\t"+value.Id)
				fmt.Fprintln(writer, "Name\t"+value.Name)
				fmt.Fprintln(writer, "Description\t"+value.Description)
				fmt.Fprintln(writer, "Command\t"+value.Command)
				fmt.Fprintln(writer, "Path\t"+value.Path)
				fmt.Fprintln(writer, "Language\t"+value.Language)
				fmt.Fprintln(writer, "OS\t"+fmt.Sprint(value.Os))
				fmt.Fprintln(writer, "Script\t"+fmt.Sprint(value.Script))
				fmt.Fprintln(writer, "Params\t"+helper.PrettyPrint(value.Params))
				fmt.Fprintln(writer, "Author\t"+value.Author)

				fmt.Fprintln(writer, "-------------")

				writer.Flush()
			}
		case "short":
			fmt.Fprintln(writer, "Name\tAuthor\tDescription\tCommand")

			for _, value := range response.QueryAnswers {
				fmt.Fprintln(writer, value.Name+"\t"+value.Author+"\t"+value.Description+"\t"+value.Command)
			}
		case "wide":
			fmt.Fprintln(writer, "Name\tAuthor\tDescription\tLanguage\tPath\tInteractive\tScript\tCommand")
			for _, value := range response.QueryAnswers {
				fmt.Fprintln(writer, value.Name+"\t"+value.Author+"\t"+value.Description+"\t"+value.Language+"\t"+value.Path+"\t"+fmt.Sprint(value.Interactive)+"\t"+fmt.Sprint(value.Script)+"\t"+value.Command)
			}
		case "json":
			jsonData, _ := json.MarshalIndent(response, "", "\t")
			fmt.Println(string(jsonData))
		default:
			fmt.Println("Undefined output")
		}

		writer.Flush()

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	searchCmd.Flags().Int32P("count", "c", 10, "Sets the maximum count of entries")
	searchCmd.Flags().StringP("output", "o", "short", "Set the output format. Options are: short, wide, json, detailed")
}
