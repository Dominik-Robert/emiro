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
	"os"
	"os/exec"
	"strings"

	"github.com/dominik-robert/emiro/emironetwork"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		emiroHost := viper.GetString("emiroHost")
		emiroPort := viper.GetInt("emiroPort")

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

		response, err := c.SendExec(context.Background(), &query)

		if err != nil {
			log.Fatalf("Error when calling SendQuery: %s", err)
		}

		parameter, _ := cmd.Flags().GetStringArray("param")

		for _, value := range parameter {
			key, value := splitKeyValue(value)
			response.Command = strings.ReplaceAll(response.Command, key, value)
		}

		if !response.Script {
			// Is no script. So execute it directly
			cmd := exec.Command(response.Path, "-c", response.Command)

			stdoutStdErr, err := cmd.CombinedOutput()

			if err != nil {
				log.Fatalf("Cannot execute command: %s", err)
			}

			fmt.Printf("%s\n", stdoutStdErr)

		} else {
			tempDir := viper.GetString("tempDir")
			file, err := ioutil.TempFile(tempDir, "emiro.*.tmp")

			if err != nil {
				log.Fatalf("Cannot create temp file %s", err)
			}

			err = file.Chmod(os.FileMode(0777))

			if err != nil {
				log.Fatalf("Cannot give execution rights to temp file %s", err)
			}

			ioutil.WriteFile(file.Name(), []byte(response.Command), os.FileMode(0777))

			cmd := exec.Command(file.Name())

			stdoutStdErr, err := cmd.CombinedOutput()

			if err != nil {
				log.Fatalf("Cannot execute command: %s", err)
			}

			fmt.Printf("%s\n", stdoutStdErr)

			os.Remove(file.Name())
		}
	},
}

func splitKeyValue(keyVal string) (key, value string) {
	arr := strings.Split(keyVal, "=")

	return arr[0], arr[1]
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	execCmd.Flags().StringArrayP("param", "p", nil, "Specify a Parameter array to change the command")
}
