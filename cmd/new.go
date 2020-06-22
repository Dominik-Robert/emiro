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
	"io/ioutil"
	"log"

	"github.com/dominik-robert/emiro/emironetwork"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		emiroHost := viper.GetString("emiroHost")
		emiroPort := viper.GetInt("emiroPort")
		conn, err := grpc.Dial(emiroHost+":"+fmt.Sprint(emiroPort), grpc.WithInsecure())

		if err != nil {
			log.Fatalf("Could not connect to server: %s", err)
		}

		c := emironetwork.NewEmiroClient(conn)

		fromFile, _ := cmd.Flags().GetString("fromFile")

		if fromFile != "" {
			// Load all from a file
			fileData, _ := ioutil.ReadFile(fromFile)

			query := emironetwork.QueryFull{
				Data: fileData,
			}

			response, err := c.SendNew(context.Background(), &query)

			if err != nil {
				log.Fatalf("Error when calling SendNew: %s", err)
			}

			if response.Succeed {
				fmt.Println("Sucessfully created")
			} else {
				fmt.Println("An error occured while creating")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	newCmd.Flags().BoolP("interactive", "i", false, "Help message for toggle")
	newCmd.Flags().StringP("name", "n", "", "Help message for toggle")
	newCmd.Flags().StringP("description", "d", "", "Help message for toggle")
	newCmd.Flags().String("fromFile", "", "Specify a file that is used")

}
