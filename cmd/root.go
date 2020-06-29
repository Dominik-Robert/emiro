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
	"os"

	"github.com/dominik-robert/emiro/config"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "emiro",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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
	cobra.OnInitialize(initConfig)

	config.Config = viper.GetViper()
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().String("databaseHost", "localhost", "Set the database host")
	rootCmd.PersistentFlags().Int("databasePort", 9200, "Set database port")
	rootCmd.PersistentFlags().String("databaseType", "elasticsearch", "Set the database Type ")
	rootCmd.PersistentFlags().String("databaseIndex", "emiro", "Sets the index which are used")
	rootCmd.PersistentFlags().Bool("databaseInsecure", false, "enable insecure database connection")
	rootCmd.PersistentFlags().String("tempDir", "/tmp/emiro", "Specifies the tempDir")
	rootCmd.PersistentFlags().String("emiroHost", "localhost", "specify the emiro host")
	rootCmd.PersistentFlags().Int("emiroPort", 9000, "specify the emiro host port")
	rootCmd.PersistentFlags().String("aliasFile", "$(HOME)/.profile", "Specify the aliasFile where the alias commands are stored to")

	viper.BindPFlag("databaseHost", rootCmd.PersistentFlags().Lookup("databaseHost"))
	viper.BindPFlag("databasePort", rootCmd.PersistentFlags().Lookup("databasePort"))
	viper.BindPFlag("databaseIndex", rootCmd.PersistentFlags().Lookup("databaseIndex"))
	viper.BindPFlag("databaseType", rootCmd.PersistentFlags().Lookup("databaseType"))
	viper.BindPFlag("databaseInsecure", rootCmd.PersistentFlags().Lookup("databaseInsecure"))
	viper.BindPFlag("tempDir", rootCmd.PersistentFlags().Lookup("tempDir"))
	viper.BindPFlag("emiroHost", rootCmd.PersistentFlags().Lookup("emiroHost"))
	viper.BindPFlag("emiroPort", rootCmd.PersistentFlags().Lookup("emiroPort"))
	viper.BindPFlag("aliasFile", rootCmd.PersistentFlags().Lookup("aliasFile"))

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/emiro.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Activates the verbose output")
}

func init() {
	// Prepare tempDir
	_ = os.Mkdir(viper.GetString("tempDir"), os.FileMode(0777))

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
		viper.SetConfigName("emiro")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	// Output loaded config if verbose output is on
	verbose, _ := rootCmd.Flags().GetBool("verbose")
	if verbose {
		err := viper.ReadInConfig()

		if err == nil {
			log.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			log.Printf("Cannot find a config file: %s", err)
		}
	}
}

// GetRootCMD Returns the cobra rootCMD for generating the documentation with cobra.
func GetRootCMD() *cobra.Command {
	return rootCmd
}
