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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Get detailed information about a command",
	Long: `With show you can get a detailed view of an existing command.
	
	For example:

	emiro show Hostname

	results in following output
	
	ID            hveX_3IB8mo66xNibt7z
	Name          echo Hostname
	Description   Prints the Hostname of the machine
	Command       echo $HOSTNAME
	Path          /bin/bash
	Language      
	OS            []
	Script        false
	Params				null`,
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"describe"},
	Run: func(cmd *cobra.Command, args []string) {
		emiroHost := viper.GetString("emiroHost")
		emiroPort := viper.GetInt("emiroPort")
		plain, _ := cmd.Flags().GetBool("plain")

		var conn *grpc.ClientConn

		conn, err := grpc.Dial(emiroHost+":"+fmt.Sprint(emiroPort), grpc.WithInsecure())

		if err != nil {
			log.Fatalf("Could not connect to server: %s", err)
		}

		c := emironetwork.NewEmiroClient(conn)

		query := emironetwork.Query{
			Query: args[0],
			All:   false,
			Count: 1,
		}

		response, err := c.SendShow(context.Background(), &query)

		if err != nil {
			log.Fatalf("Error when calling SendQuery: %s", err)
		}

		if plain {
			result, _ := json.Marshal(response)
			fmt.Println(string(result))
			return
		}

		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(writer, "ID\t"+response.Id)
		fmt.Fprintln(writer, "Name\t"+response.Name)
		fmt.Fprintln(writer, "Description\t"+response.Description)
		fmt.Fprintln(writer, "Command\t"+response.Command)
		fmt.Fprintln(writer, "Path\t"+response.Path)
		fmt.Fprintln(writer, "Language\t"+response.Language)
		fmt.Fprintln(writer, "OS\t"+fmt.Sprint(response.Os))
		fmt.Fprintln(writer, "Script\t"+fmt.Sprint(response.Script))
		fmt.Fprintln(writer, "Params\t"+PrettyPrint(response.Params))

		writer.Flush()
	},
}

func PrettyPrint(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	showCmd.Flags().BoolP("plain", "p", false, "Shows the result in plain json")

}
