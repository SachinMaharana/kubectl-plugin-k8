package cmd

import (
	"fmt"
	"os"

	"github.com/sachin/kubectl-plugin/pkg/capacity"
	"github.com/spf13/cobra"
)

var showPods bool

var rootCmd = &cobra.Command{
	Use:   "kubectl-plugin",
	Short: "test",
	Long:  "test",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.ParseFlags(args); err != nil {
			fmt.Printf("Error parsing flags: %v", err)
		}
		capacity.List(args, showPods)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `Version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo")
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&showPods, "pods", "p", false, "Set this flag to include pods in output")
}

//Execute ...
func Execute() {
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
