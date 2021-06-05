package cmd

import (
	"fmt"
	"github.com/smuthusa/taskexecutor/src/handlers"
	"github.com/smuthusa/taskexecutor/src/task"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:  "execute",
	Long: "execute",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			cmd.Usage()
			log.Fatalf("missing file path")
		}

		if _, err := os.Stat(file); err != nil {
			log.Fatalf("missing file %s!", file)
		}
		validate, _ := cmd.Flags().GetBool("validate")
		if validate {
			task.Execute(file, handlers.ConsoleLogger)
		} else {
			task.Execute(file, handlers.CommandExecutor)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("file", "f", "<file path>", "Task definition file path")
	rootCmd.PersistentFlags().Bool("validate", false, "Do not execute, validate the command flow")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
