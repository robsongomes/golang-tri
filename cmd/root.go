/*
Copyright © 2022 Robson Gomes

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dataFile string
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "tri is a Todo application",
	Long: `tri will help you get more done in less time.
It's designed to be as simple as possible to help
you accomplish your goals.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list`,
	Run:   addRun,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	if len(dataFile) > 0 {
		viper.Set("datafile", dataFile)
	} else {
		if len(cfgFile) > 0 {
			viper.SetConfigFile(cfgFile)
		} else {
			viper.SetConfigName("tri")
			viper.SetConfigType("yaml")
			viper.AddConfigPath(".")
			viper.AddConfigPath("$HOME")
			viper.AutomaticEnv()
			viper.SetEnvPrefix("tri")
		}

		err := viper.ReadInConfig()

		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		} else {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}

}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/tri.yaml)")
	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", "", "data file to store todos")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(addCmd)
}
