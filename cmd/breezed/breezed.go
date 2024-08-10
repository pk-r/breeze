package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pk-r/breeze-agent/pkg/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Version string

var rootCmd = &cobra.Command{
	Use:   "breezed",
	Short: "breeze daemon worker node",
	Long:  `This process waits for commands to initiate a job`,
	Run: func(cmd *cobra.Command, args []string) {
		s := storage.NewGitStorage("https://github.com/go-git/go-billy")
		files, _ := s.FetchFiles(context.Background(), "hi")
		fmt.Println(string(files[0]))
	},
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number of breezed",
	Long:  `This command allows you to print the version number of breezed.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", Version)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(versionCmd)
}

func initConfig() {
	viper.SetEnvPrefix("BREEZED")
	viper.AutomaticEnv()

	// Read in environment variables from the .env file
	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
