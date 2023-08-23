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
	"log"
	"net"

	"github.com/dominik-robert/emiro/config"
	"github.com/dominik-robert/emiro/database"
	"github.com/dominik-robert/emiro/emironetwork"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts an emiro server",
	Long: `With server you simply start a new emiro server where you can connect from your client.
Example:

emiro server

You can change the database connection with a config file or parameter. You can even change the listening port for the server with --emiroServerHost and --emiroServerPort `,
	Run: func(cmd *cobra.Command, args []string) {
		emiroHost, _ := rootCmd.Flags().GetString("emiroServerHost")

		if emiroHost == "" {
			emiroHost = config.ConfigStruct.Server.Hostname
		}

		emiroPort, _ := rootCmd.Flags().GetInt("emiroServerPort")
		if emiroPort == 9000 {
			emiroPort = config.ConfigStruct.Server.Port
		}
		databaseType, _ := cmd.Flags().GetString("type")

		lis, err := net.Listen("tcp", emiroHost+fmt.Sprint(emiroPort))

		if err != nil {
			log.Fatal(err)
		}
		defer lis.Close()
		fmt.Println("Loaded database type", databaseType)
		database.DB = database.InitDatabase(databaseType)
		log.Println("Connect to Database")
		log.Println("Started Server on Port " + fmt.Sprint(emiroPort))

		s := emironetwork.Server{}

		grpcServer := grpc.NewServer()

		emironetwork.RegisterEmiroServer(grpcServer, &s)
		log.Println("Started gRPC Server")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server over port 9000: %s", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serverCmd.PersistentFlags().String("interface", "", "Binds the server on the specified interface")
	serverCmd.PersistentFlags().String("type", "sqlite", "The type of database. Supported is sqlite and elasticsearch")
	serverCmd.PersistentFlags().Int("listenPort", 9000, "Binds the server on the specified Port")
}
