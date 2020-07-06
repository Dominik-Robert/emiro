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
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/dominik-robert/emiro/alias"
	"github.com/dominik-robert/emiro/emironetwork"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a snippet",
	Long: `Executes a snippet from the terminal. You can write a query which results in one command (this will automatically executed) 
or you can refer to an id with --id where you don't have to write a more detailed query.

For example:

emiro exec "echo hostname"

# Will execute the query and execute the resulting command.

You can execute your snippet on a remote machine with --remote HOST:PORT.

You can append or prepend arguments to your command with --prepend or --append, if you write more arguments than one (first is the query) 
then the others will automatically appended to the command.`,
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"run"},
	Run: func(cmd *cobra.Command, args []string) {
		aliasFlag, _ := cmd.Flags().GetString("alias")
		remote, _ := cmd.Flags().GetString("remote")
		parameter, _ := cmd.Flags().GetStringArray("params")
		appendVar, _ := cmd.Flags().GetString("append")
		prependVar, _ := cmd.Flags().GetString("prepend")

		if aliasFlag != "" {
			aliasFilePath := viper.GetString("aliasFile")
			home, _ := homedir.Dir()
			aliasFilePath = strings.ReplaceAll(aliasFilePath, "$(HOME)", home)
			aliasFile := alias.NewAlias(aliasFilePath)

			reg := regexp.MustCompile(" --alias \\w+")
			programCall := reg.ReplaceAllLiteralString(strings.Join(os.Args, " "), "")
			err := aliasFile.UpdateAliasFile(aliasFlag, programCall)

			if err != nil {
				log.Fatalf("Problem while adding alias: %s", err)
			}
		}

		emiroHost := viper.GetString("emiroHost")
		emiroPort := viper.GetInt("emiroPort")
		verbose, _ := rootCmd.Flags().GetBool("verbose")

		var conn *grpc.ClientConn

		conn, err := grpc.Dial(emiroHost+":"+fmt.Sprint(emiroPort), grpc.WithInsecure())

		if err != nil {
			log.Fatalf("Could not connect to server: %s", err)
		}

		c := emironetwork.NewEmiroClient(conn)

		if len(args) > 1 {
			appendVar = strings.Join(args[1:], " ") + " " + appendVar
		}

		query := emironetwork.Query{
			Query:      args[0],
			All:        false,
			Count:      1,
			RemoteHost: remote,
			Parameter:  parameter,
			Append:     appendVar,
			Prepend:    prependVar,
		}

		if remote != "" {
			streamResponse, err := c.ExecRemote(context.Background(), &query)
			if err != nil {
				log.Println(err)
			}

			responseTmp, err := streamResponse.Recv()

			if err != nil {
				log.Println(err)
			}

			for responseTmp.Succeed {
				fmt.Println(string(responseTmp.Data))
				responseTmp, err = streamResponse.Recv()

				if err != nil {
					return
				}
			}

			return
		}

		response, err := c.SendExec(context.Background(), &query)

		if err != nil {
			log.Fatalf("Error when calling SendQuery: %s", err)
		}

		if verbose {
			log.Println("Running Command:", response.Command)
		}

		if !response.Script {
			// Is no script. So execute it directly
			cmd := exec.Command(response.Path, "-c", response.Command)

			stdout, errOut := cmd.StdoutPipe()
			stdErr, errErr := cmd.StderrPipe()
			cmd.Stdin = os.Stdin

			if errOut != nil {
				log.Fatalf("Cannot connect to commands stdOut: %s", errOut)
			}
			if errErr != nil {
				log.Fatalf("Cannot connect to commands stdErr: %s", errErr)
			}

			cmd.Start()

			go func() {
				scannerErr := bufio.NewScanner(stdErr)
				for scannerErr.Scan() {
					m := scannerErr.Text()
					fmt.Println(m)
				}
			}()

			scannerOut := bufio.NewScanner(stdout)
			for scannerOut.Scan() {
				m := scannerOut.Text()
				fmt.Println(m)
			}
			cmd.Wait()

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

			stdout, errOut := cmd.StdoutPipe()
			stdErr, errErr := cmd.StderrPipe()

			if errOut != nil {
				log.Fatalf("Cannot connect to commands stdOut: %s", err)
			}
			if errErr != nil {
				log.Fatalf("Cannot connect to commands stdErr: %s", err)
			}

			cmd.Start()

			go func() {
				scannerErr := bufio.NewScanner(stdErr)
				for scannerErr.Scan() {
					m := scannerErr.Text()
					fmt.Println(m)
				}
			}()

			scannerOut := bufio.NewScanner(stdout)
			for scannerOut.Scan() {
				m := scannerOut.Text()
				fmt.Println(m)
			}
			cmd.Wait()

			os.Remove(file.Name())
		}
	},
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
	execCmd.Flags().String("alias", "", "Specify if the command will create an alias in your system")
	execCmd.Flags().String("remote", "", "Specify the host where the command will run")
	execCmd.Flags().String("append", "", "Append the argument after the command")
	execCmd.Flags().String("prepend", "", "Prepend the argument before the command")

}
