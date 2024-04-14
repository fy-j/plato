package cmd

import "github.com/spf13/cobra"

func init() {
	cobra.OnInitialize(initConfig)
}

var rootCmd = &cobra.Command{
	Use:   "plato",
	Short: "im server",
	Run:   plato,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func plato(cmd *cobra.Command, args []string) {

}

func initConfig() {

}
