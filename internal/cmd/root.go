package cmd

import (
	"fmt"
	"os"

	cube "cubectl/internal/app/cube"
	"cubectl/internal/cmd/delete"
	"cubectl/internal/cmd/describe"
	"cubectl/internal/cmd/get"
	"cubectl/internal/cmd/logs"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cubectl",
	Short: "cubectl controls cube instead of Kubernetes clusters.",
	Long: `cubectl controls cube instead of Kubernetes clusters.

Find more information at:
  https://github.com/yoshihicode/cubectl
  
Controls:
  Arrow keys or wasd: Rotate the cube
  z: Zoom in
  x: Zoom out
  Ctrl+C or Esc: Exit`,
	Run: func(cmd *cobra.Command, args []string) {
		output, _ := cmd.Flags().GetString("output")
		watch, _ := cmd.Flags().GetBool("watch")

		opts := cube.Options{
			Output: output,
			Watch:  watch,
		}
		// default action
		cube.Render(cmd.Context(), opts)
	},
	SilenceUsage: true,
}

func init() {
	rootCmd.SetHelpTemplate(cubectlHelpTemplate)
	rootCmd.SetUsageTemplate(cubectlUsageTemplate)

	rootCmd.AddGroup(
		&cobra.Group{
			ID:    "basic",
			Title: "Basic Cube Commands (Beginner):",
		},
		&cobra.Group{
			ID:    "troubleshooting",
			Title: "Troubleshooting and Debugging Commands:",
		},
	)

	rootCmd.AddCommand(describe.NewDescribeCmd())
	rootCmd.AddCommand(get.NewGetCmd())
	rootCmd.AddCommand(logs.NewLogsCmd())
	rootCmd.AddCommand(delete.NewDeleteCmd())
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
