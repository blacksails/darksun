package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "darksun",
	Short: "🌓 switch your applications between dark and sun mode",
}

var darkCmd = &cobra.Command{
	Use:   "dark",
	Short: "🌑 switch all configured modules to dark mode",
	Run: func(cmd *cobra.Command, args []string) {
		getAndRunModules(true)
		if err := setCurrentMode(true); err != nil {
			errExit(err)
		}
	},
}

var lightCmd = &cobra.Command{
	Use:     "sun",
	Short:   "🌕 switch all configured modules to sun mode",
	Aliases: []string{"light"},
	Run: func(cmd *cobra.Command, args []string) {
		getAndRunModules(false)
		if err := setCurrentMode(false); err != nil {
			errExit(err)
		}
	},
}

var toggleCmd = &cobra.Command{
	Use:     "toggle",
	Short:   "🌓 toggle your applications between dark and sun mode",
	Aliases: []string{"light"},
	Run: func(cmd *cobra.Command, args []string) {
		dark := cfg.GetBool("dark")
		var err error
		if dark {
			getAndRunModules(false)
			err = setCurrentMode(false)
		} else {
			getAndRunModules(true)
			err = setCurrentMode(true)
		}
		if err != nil {
			errExit(err)
		}
	},
}

func errExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func setCurrentMode(dark bool) error {
	cfg.Set("dark", dark)
	return cfg.WriteConfig()
}

func getAndRunModules(dark bool) {
	mods, err := GetModules()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	runModules(mods, dark)
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
	rootCmd.AddCommand(darkCmd, lightCmd, toggleCmd)
	rootCmd.Execute()
}
