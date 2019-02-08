package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/SachinMaharana/kubectl-plugin-k8/pkg/nodepodlist"
	"github.com/spf13/cobra"
)

var showNodePods bool

var rootCmd = &cobra.Command{
	Use:   "kubectl-pn",
	Short: "Custom Plugins",
	Long:  "https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/",
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
var installNodeWebCmd = &cobra.Command{
	Use:   "node-web",
	Short: "Install my k8s app.",
	Long:  `https://github.com/SachinMaharana/k8-port`,
	Run: func(cmd *cobra.Command, args []string) {
		_, lookErr := exec.LookPath("git")
		if lookErr != nil {
			panic(lookErr)
		}

		cmnd := exec.Command("git", "clone", "https://github.com/SachinMaharana/k8-port.git")
		log.Printf("Cloning and waiting for it to finish...")
		err := cmnd.Run()
		if err != nil {
			log.Printf("Command finished with error: %v", err)
			os.Exit(1)
		}
		fmt.Println("Deploying k8s app: node-web")
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&showNodePods, "pods-node list", "p", false, "Set this flag to include pods-node in output")
}

func main() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(installNodeWebCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
