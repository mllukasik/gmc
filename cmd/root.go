package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/mllukasik/gmc/build"
	"github.com/mllukasik/gmc/cmd/branch"
	"github.com/mllukasik/gmc/cmd/push"
	"github.com/spf13/cobra"
)

func longDesc() string {
	var builder strings.Builder
	builder.WriteString("Git Minimized Command\n")
	builder.WriteString(fmt.Sprintf("Version: %s\n", build.Build().Version))
	builder.WriteString(fmt.Sprintf("Build date: %s\n", build.Build().Date.Format(time.ANSIC)))
	return builder.String()
}

var rootCmd = &cobra.Command{
	Use:   "gmc",
	Short: "Git Minimized Command",
	Long:  longDesc(),
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetBool("version")
		if (version) {
			fmt.Println(build.Build().Version)
			return
		}
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(branch.BranchCmd)
	rootCmd.AddCommand(push.PushCmd)
	rootCmd.PersistentFlags().Bool("version", false, "Shows cli version")
}

