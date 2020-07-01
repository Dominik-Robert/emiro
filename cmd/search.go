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
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/dominik-robert/emiro/emironetwork"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
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
	emiro search kubernetes -c 5
	
	`,
	Aliases: []string{"list"},
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		count, _ := cmd.Flags().GetInt32("count")
		emiroHost := viper.GetString("emiroHost")
		emiroPort := viper.GetInt("emiroPort")

		var conn *grpc.ClientConn

		conn, err := grpc.Dial(emiroHost+":"+fmt.Sprint(emiroPort), grpc.WithInsecure())

		if err != nil {
			log.Fatalf("Could not connect to server: %s", err)
		}

		c := emironetwork.NewEmiroClient(conn)

		queryString := ""

		if !all {
			if len(args) != 1 {
				log.Fatalf("Needs one argument or -a Parameter")
			}
			queryString = args[0]
		}

		query := emironetwork.Query{
			Query: queryString,
			All:   all,
			Count: count,
		}

		response, err := c.SendQuery(context.Background(), &query)

		if err != nil {
			log.Fatalf("Error when calling SendQuery: %s", err)
		}

		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(writer, "Name\tDescription\tLanguage\tPath\tInteractive\tScript")

		for _, value := range response.QueryAnswers {
			fmt.Fprintln(writer, value.Name+"\t"+value.Description+"\t"+value.Language+"\t"+value.Path+"\t"+fmt.Sprint(value.Interactive)+"\t"+fmt.Sprint(value.Script))
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
	searchCmd.Flags().BoolP("all", "a", false, "Shows all existing entries")
	searchCmd.Flags().Int32P("count", "c", 10, "Sets the maximum count of entries")
}
