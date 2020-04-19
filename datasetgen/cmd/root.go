package cmd

import (
	"fmt"
	"os"

	datasetgen "github.com/covid19-aomori/go-datasetgen"
	"github.com/spf13/cobra"
)

var Version string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "datasetgen",
	Short:        "",
	Version:      Version,
	RunE:         datasetgen.Run,
	SilenceUsage: true,
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
}

func initConfig() {}
