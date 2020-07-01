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
	"log"
	"net"

	"github.com/dominik-robert/emiro/emironetwork"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts an emiro server",
	Long: `With server you simply start a new emiro server where you can connect from your client.
	
	Example:
	
	emiro server`,
	Run: func(cmd *cobra.Command, args []string) {
		emironetwork.DoCurl(viper.GetString("databaseHost"), viper.GetInt("databasePort"), "PUT", viper.GetString("databaseIndex"), "", viper.GetBool("databaseInsecure"), nil)

		lis, err := net.Listen("tcp", ":9000")

		if err != nil {
			log.Fatal(err)
		}
		log.Println("Started Server on Port 9000")

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
	serverCmd.PersistentFlags().Int("listenPort", 9000, "Binds the server on the specified Port")
}
