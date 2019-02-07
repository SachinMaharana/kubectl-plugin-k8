package cmd

import (
	"fmt"
	"os"

	"github.com/sachin/kubectl-plugin-k8/pkg/nodepodlist"
	"github.com/spf13/cobra"
)

var showNodePods bool

var rootCmd = &cobra.Command{
	Use:   "kubectl-pn",
	Short: "test",
	Long:  "test",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.ParseFlags(args); err != nil {
			fmt.Printf("Error parsing flags: %v", err)
		}
		nodepodlist.List(args, showNodePods)
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of this plugin.",
	Long:  `Version of kubectl-pn`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.1.0")
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&showNodePods, "pods-node list", "p", false, "Set this flag to include pods-node in output")
}

//Execute ...
func Execute() {
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
