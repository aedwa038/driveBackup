package main

import (
	"flag"

	"github.com/aedwa038/driveBackup/cmd"
	"github.com/spf13/cobra"
)

var folder = flag.String("folder", "Zencastr", "url input")

var rootCmd = &cobra.Command{Use: "app"}

func main() {

	cmd.Execute()

}
