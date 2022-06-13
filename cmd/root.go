package cmd

import (
	"fmt"
	"os"

	"github.com/cake-cutter/cli/cli/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cakecutter",
	Short: "Create projects from pre-built cakes (templates)! Supports files, packages, content, running commands and more!",
	Long:  `Cakecutter is a tool for creating projects from pre-built cakes (templates)! Supports files, packages, content, running commands and more!`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(utils.Colorize("red", "Error: "+err.Error()))
		os.Exit(1)
	}
}
