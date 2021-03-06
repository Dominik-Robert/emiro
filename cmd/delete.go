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

	"github.com/dominik-robert/emiro/emironetwork"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes no more needed commands",
	Long: `Deletes all commands that are matching the specified argument. 

Syntax be like:

emiro delete "SSH"
emiro delete "Mutliline String"

This will delete all Commands that have the queries (SSH or "Multiline String") in the name.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")

		count := int32(1)
		if all {
			count = 100
		}

		emiroHost := viper.GetString("emiroHost")
		emiroPort := viper.GetInt("emiroPort")
		conn, err := grpc.Dial(emiroHost+":"+fmt.Sprint(emiroPort), grpc.WithInsecure())
		defer conn.Close()
		if err != nil {
			log.Fatalf("Could not connect to server: %s", err)
		}

		c := emironetwork.NewEmiroClient(conn)

		query := emironetwork.Query{
			Query: args[0],
			All:   all,
			Count: count,
		}

		response, err := c.SendDelete(context.Background(), &query)

		if response.Succeed {
			fmt.Println("Sucessfully deleted")
		} else {
			fmt.Println("An error occured while deleting")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	deleteCmd.Flags().BoolP("all", "a", false, "Delete all Matching. If false deletes only the first matching")
}
