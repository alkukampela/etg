package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/alkukampela/etg/charconv"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "etg",
	Short: "A brief description of etg",
	Long: `A longer description that plaas multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		message := strings.Join(args, " ")

		zeta := charconv.ToBoolMap(message)

		printZeta(zeta)
	},
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
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func printZeta(zeta [][]bool) {
	for y, row := range zeta {
		for x, _ := range row {
			if zeta[y][x] {
				fmt.Print("▓")
			} else {
				fmt.Print("░")
			}
		}
		fmt.Println("")
	}
}
