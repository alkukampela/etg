package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/alkukampela/etg/charconv"

	"github.com/spf13/cobra"
)

//var nonResponsive bool

var rootCmd = &cobra.Command{
	Use:   "etg",
	Short: "A brief description of etg",
	Long: `A longer description that plaas multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		message := strings.Join(args, " ")
		boolMap := charconv.MessageToBoolMap(message)

		printBoolMap(boolMap)
	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("nonresponsive", "n", false, "print message to non responsive mode (in one line)")
//	rootCmd.Flags().Bool(&nonResponsive, "nonresponsive", "n",
//		"Print message to non responsive mode (in one line)")
}


func printBoolMap(boolMap [][]bool) {
	for y, row := range boolMap {
		for x, _ := range row {
			if boolMap[y][x] {
				fmt.Print("▓")
			} else {
				fmt.Print("░")
			}
		}
		fmt.Println("")
	}
}
