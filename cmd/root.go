package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/alkukampela/etg/charconv"
	"github.com/alkukampela/etg/messageformatter"

	"github.com/spf13/cobra"
)

var (
	responsive bool

	rootCmd = &cobra.Command{
		Use:   "etg",
		Short: "A brief description of etg",
		Long: `Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			message := strings.Join(args, " ")
			formattedMessage := messageformatter.Format(message, responsive)
			boolMap := charconv.MessageToBoolMap(formattedMessage)

			printBoolMap(boolMap)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&responsive, "responsive", "r", false,
		"print message in responsive mode (one letter per line)")
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
