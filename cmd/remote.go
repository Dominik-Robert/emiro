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

var (
	accessMap map[string]string
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Runs a remote server",
	Long: `A remote server is needed when you want to execute a command on a different machine than localhost.
You can do it with:

emiro remote`,
	Run: func(cmd *cobra.Command, args []string) {
		emironetwork.DoCurl(viper.GetString("databaseHost"), viper.GetInt("databasePort"), "PUT", viper.GetString("databaseIndex"), "", viper.GetBool("databaseInsecure"), nil)

		lis, err := net.Listen("tcp", ":9001")
		defer lis.Close()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Started Server on Port 9001")

		s := emironetwork.Server{}

		grpcServer := grpc.NewServer()

		emironetwork.RegisterEmiroServer(grpcServer, &s)

		log.Println("Started gRPC remote Server")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server over port 9001: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
