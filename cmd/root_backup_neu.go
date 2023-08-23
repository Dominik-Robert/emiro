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
	"io/ioutil"
	"log"
	"os"

	"github.com/dominik-robert/emiro/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// VERSION shows the current version of the program
var VERSION = "0.9.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "emiro",
	Short: "Emiro - Search, find and execute shell snippets without leaving your terminal",
	Long: `Emiro was build for the problem that you need to leave your terminal for every snippet you want to insert in your terminal. 
With Emiro you can search for specific snippets and execute them. 
If you don't have the needed dependencies you can even run it everywhere(where you ran emiro remote service is running.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	config.InitConfig()

	// cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().String("databaseHost", "http://localhost", "Set the database host")
	rootCmd.PersistentFlags().Int("databasePort", 9200, "Set database port")
	rootCmd.PersistentFlags().String("databaseType", "elasticsearch", "Set the database Type ")
	rootCmd.PersistentFlags().String("databaseIndex", "emiro", "Sets the index which are used")
	rootCmd.PersistentFlags().Bool("databaseInsecure", false, "enable insecure database connection")
	rootCmd.PersistentFlags().StringP("tempDir", "", "/tmp/emiro", "Specifies the tempDir")
	rootCmd.PersistentFlags().StringP("emiroHost", "", "104.197.10.159", "specify the emiro host")
	rootCmd.PersistentFlags().Int("emiroPort", 9000, "specify the emiro host port")
	rootCmd.PersistentFlags().String("aliasFile", "$(HOME)/.profile", "Specify the aliasFile where the alias commands are stored to")
	rootCmd.PersistentFlags().String("emiroServerHost", ":", "Specify the listening host for the server")
	rootCmd.PersistentFlags().Int("emiroServerPort", 9000, "Specify the listening port for the server")

	viper.BindPFlag("databaseHost", rootCmd.PersistentFlags().Lookup("databaseHost"))
	viper.BindPFlag("databasePort", rootCmd.PersistentFlags().Lookup("databasePort"))
	viper.BindPFlag("databaseIndex", rootCmd.PersistentFlags().Lookup("databaseIndex"))
	viper.BindPFlag("databaseType", rootCmd.PersistentFlags().Lookup("databaseType"))
	viper.BindPFlag("databaseInsecure", rootCmd.PersistentFlags().Lookup("databaseInsecure"))
	viper.BindPFlag("tempDir", rootCmd.PersistentFlags().Lookup("tempDir"))
	viper.BindPFlag("emiroHost", rootCmd.PersistentFlags().Lookup("emiroHost"))
	viper.BindPFlag("emiroPort", rootCmd.PersistentFlags().Lookup("emiroPort"))
	viper.BindPFlag("aliasFile", rootCmd.PersistentFlags().Lookup("aliasFile"))
	viper.BindPFlag("emiroServerHost", rootCmd.PersistentFlags().Lookup("emiroServerHost"))
	viper.BindPFlag("emiroServerPort", rootCmd.PersistentFlags().Lookup("emiroServerPort"))

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/emiro.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Activates the verbose output")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".emiro" (without extension).
		viper.AddConfigPath("./")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/emiro/")
		viper.SetConfigName(".emiro")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	// Output loaded config if verbose output is on
	verbose, _ := rootCmd.Flags().GetBool("verbose")
	err := viper.ReadInConfig()
	if verbose {
		if err == nil {
			log.Println("Using config file:", viper.ConfigFileUsed())
			cfgFile = viper.ConfigFileUsed()
		} else {
			log.Printf("Cannot find a config file: %s", err)
		}
	}

	data, err := ioutil.ReadFile(viper.ConfigFileUsed())
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal(data, &config.CfgStruct)

	if err != nil {
		log.Println(err)
	}
}

func init() {
	// Prepare tempDir
	_ = os.Mkdir(viper.GetString("tempDir"), os.FileMode(0777))
}

// GetRootCMD Returns the cobra rootCMD for generating the documentation with cobra.
func GetRootCMD() *cobra.Command {
	return rootCmd
}
