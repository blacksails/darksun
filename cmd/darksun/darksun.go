package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "darksun",
	/*
		Run: func(cmd *cobra.Command, args []string) {
			m := &iterm.Module{}
			err := m.Light()
			fmt.Println(err)
		},*/
}

var darkCmd = &cobra.Command{
	Use: "dark",
	Run: func(cmd *cobra.Command, args []string) {
		mods, err := GetModules()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		runModules(mods, true)
	},
}

var lightCmd = &cobra.Command{
	Use: "sun",
	Run: func(cmd *cobra.Command, args []string) {
		mods, err := GetModules()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		runModules(mods, false)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

var cfg *viper.Viper

func initConfig() {
	cfg = viper.New()
	cfg.SetConfigName("config")
	cfg.AddConfigPath("$HOME/.config/darksun/")
	cfg.AddConfigPath("$HOME/.darksun/")
	cfg.AddConfigPath("/etc/darksun/")
	cfg.AddConfigPath(".")
	if err := cfg.ReadInConfig(); err != nil {
		fmt.Println("Could not read in config:", err)
		os.Exit(1)
	}
}

func main() {
	rootCmd.AddCommand(darkCmd, lightCmd)
	rootCmd.Execute()
}
